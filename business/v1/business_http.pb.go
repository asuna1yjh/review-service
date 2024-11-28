// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             v3.20.1
// source: business/v1/business.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationBusinessAppealReview = "/api.business.v1.Business/AppealReview"
const OperationBusinessReplyReview = "/api.business.v1.Business/ReplyReview"

type BusinessHTTPServer interface {
	// AppealReview 商家申诉用户评价
	AppealReview(context.Context, *AppealReviewRequest) (*AppealReviewReply, error)
	// ReplyReview 商家回复用户评价
	ReplyReview(context.Context, *ReplyReviewRequest) (*ReplyReviewReply, error)
}

func RegisterBusinessHTTPServer(s *http.Server, srv BusinessHTTPServer) {
	r := s.Route("/")
	r.POST("b/v1/review/reply", _Business_ReplyReview0_HTTP_Handler(srv))
	r.POST("b/v1/review/appeal", _Business_AppealReview0_HTTP_Handler(srv))
}

func _Business_ReplyReview0_HTTP_Handler(srv BusinessHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReplyReviewRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBusinessReplyReview)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ReplyReview(ctx, req.(*ReplyReviewRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ReplyReviewReply)
		return ctx.Result(200, reply)
	}
}

func _Business_AppealReview0_HTTP_Handler(srv BusinessHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AppealReviewRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBusinessAppealReview)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AppealReview(ctx, req.(*AppealReviewRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AppealReviewReply)
		return ctx.Result(200, reply)
	}
}

type BusinessHTTPClient interface {
	AppealReview(ctx context.Context, req *AppealReviewRequest, opts ...http.CallOption) (rsp *AppealReviewReply, err error)
	ReplyReview(ctx context.Context, req *ReplyReviewRequest, opts ...http.CallOption) (rsp *ReplyReviewReply, err error)
}

type BusinessHTTPClientImpl struct {
	cc *http.Client
}

func NewBusinessHTTPClient(client *http.Client) BusinessHTTPClient {
	return &BusinessHTTPClientImpl{client}
}

func (c *BusinessHTTPClientImpl) AppealReview(ctx context.Context, in *AppealReviewRequest, opts ...http.CallOption) (*AppealReviewReply, error) {
	var out AppealReviewReply
	pattern := "b/v1/review/appeal"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBusinessAppealReview))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BusinessHTTPClientImpl) ReplyReview(ctx context.Context, in *ReplyReviewRequest, opts ...http.CallOption) (*ReplyReviewReply, error) {
	var out ReplyReviewReply
	pattern := "b/v1/review/reply"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBusinessReplyReview))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
