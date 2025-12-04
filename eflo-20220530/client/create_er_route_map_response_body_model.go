// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateErRouteMapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateErRouteMapResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *CreateErRouteMapResponseBody
	GetCode() *int32
	SetContent(v *CreateErRouteMapResponseBodyContent) *CreateErRouteMapResponseBody
	GetContent() *CreateErRouteMapResponseBodyContent
	SetMessage(v string) *CreateErRouteMapResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateErRouteMapResponseBody
	GetRequestId() *string
}

type CreateErRouteMapResponseBody struct {
	// The details about the access denial. This parameter is returned only if Resource Access Management (RAM) authentication failed.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *CreateErRouteMapResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 039C3C3A-3C37-5672-80D5-D8CD48C676D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateErRouteMapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateErRouteMapResponseBody) GoString() string {
	return s.String()
}

func (s *CreateErRouteMapResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateErRouteMapResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateErRouteMapResponseBody) GetContent() *CreateErRouteMapResponseBodyContent {
	return s.Content
}

func (s *CreateErRouteMapResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateErRouteMapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateErRouteMapResponseBody) SetAccessDeniedDetail(v string) *CreateErRouteMapResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateErRouteMapResponseBody) SetCode(v int32) *CreateErRouteMapResponseBody {
	s.Code = &v
	return s
}

func (s *CreateErRouteMapResponseBody) SetContent(v *CreateErRouteMapResponseBodyContent) *CreateErRouteMapResponseBody {
	s.Content = v
	return s
}

func (s *CreateErRouteMapResponseBody) SetMessage(v string) *CreateErRouteMapResponseBody {
	s.Message = &v
	return s
}

func (s *CreateErRouteMapResponseBody) SetRequestId(v string) *CreateErRouteMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateErRouteMapResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateErRouteMapResponseBodyContent struct {
	// routing policy ID
	//
	// example:
	//
	// er-rmap-uwglhzom
	ErRouteMapId *string `json:"ErRouteMapId,omitempty" xml:"ErRouteMapId,omitempty"`
}

func (s CreateErRouteMapResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s CreateErRouteMapResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CreateErRouteMapResponseBodyContent) GetErRouteMapId() *string {
	return s.ErRouteMapId
}

func (s *CreateErRouteMapResponseBodyContent) SetErRouteMapId(v string) *CreateErRouteMapResponseBodyContent {
	s.ErRouteMapId = &v
	return s
}

func (s *CreateErRouteMapResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
