// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVccRouteEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateVccRouteEntryResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *CreateVccRouteEntryResponseBody
	GetCode() *int32
	SetContent(v *CreateVccRouteEntryResponseBodyContent) *CreateVccRouteEntryResponseBody
	GetContent() *CreateVccRouteEntryResponseBodyContent
	SetMessage(v string) *CreateVccRouteEntryResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateVccRouteEntryResponseBody
	GetRequestId() *string
}

type CreateVccRouteEntryResponseBody struct {
	// The detailed information about the failed permission verification.
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
	// The returned data.
	Content *CreateVccRouteEntryResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The returned message.
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

func (s CreateVccRouteEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVccRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVccRouteEntryResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateVccRouteEntryResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateVccRouteEntryResponseBody) GetContent() *CreateVccRouteEntryResponseBodyContent {
	return s.Content
}

func (s *CreateVccRouteEntryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateVccRouteEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVccRouteEntryResponseBody) SetAccessDeniedDetail(v string) *CreateVccRouteEntryResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateVccRouteEntryResponseBody) SetCode(v int32) *CreateVccRouteEntryResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVccRouteEntryResponseBody) SetContent(v *CreateVccRouteEntryResponseBodyContent) *CreateVccRouteEntryResponseBody {
	s.Content = v
	return s
}

func (s *CreateVccRouteEntryResponseBody) SetMessage(v string) *CreateVccRouteEntryResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVccRouteEntryResponseBody) SetRequestId(v string) *CreateVccRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVccRouteEntryResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateVccRouteEntryResponseBodyContent struct {
	// The ID of the route entry.
	//
	// example:
	//
	// vcc-rte-5cey1sap
	VccRouteEntryId *string `json:"VccRouteEntryId,omitempty" xml:"VccRouteEntryId,omitempty"`
}

func (s CreateVccRouteEntryResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s CreateVccRouteEntryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CreateVccRouteEntryResponseBodyContent) GetVccRouteEntryId() *string {
	return s.VccRouteEntryId
}

func (s *CreateVccRouteEntryResponseBodyContent) SetVccRouteEntryId(v string) *CreateVccRouteEntryResponseBodyContent {
	s.VccRouteEntryId = &v
	return s
}

func (s *CreateVccRouteEntryResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
