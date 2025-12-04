// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnAssociateVpdCidrBlockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UnAssociateVpdCidrBlockResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *UnAssociateVpdCidrBlockResponseBody
	GetCode() *int32
	SetContent(v *UnAssociateVpdCidrBlockResponseBodyContent) *UnAssociateVpdCidrBlockResponseBody
	GetContent() *UnAssociateVpdCidrBlockResponseBodyContent
	SetMessage(v string) *UnAssociateVpdCidrBlockResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnAssociateVpdCidrBlockResponseBody
	GetRequestId() *string
}

type UnAssociateVpdCidrBlockResponseBody struct {
	// 访问被拒绝详细信息。
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
	// The response data.
	Content *UnAssociateVpdCidrBlockResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 9C50C9CD-E799-54DA-BA7A-1FAF3DF80857
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnAssociateVpdCidrBlockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnAssociateVpdCidrBlockResponseBody) GoString() string {
	return s.String()
}

func (s *UnAssociateVpdCidrBlockResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UnAssociateVpdCidrBlockResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UnAssociateVpdCidrBlockResponseBody) GetContent() *UnAssociateVpdCidrBlockResponseBodyContent {
	return s.Content
}

func (s *UnAssociateVpdCidrBlockResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnAssociateVpdCidrBlockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnAssociateVpdCidrBlockResponseBody) SetAccessDeniedDetail(v string) *UnAssociateVpdCidrBlockResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UnAssociateVpdCidrBlockResponseBody) SetCode(v int32) *UnAssociateVpdCidrBlockResponseBody {
	s.Code = &v
	return s
}

func (s *UnAssociateVpdCidrBlockResponseBody) SetContent(v *UnAssociateVpdCidrBlockResponseBodyContent) *UnAssociateVpdCidrBlockResponseBody {
	s.Content = v
	return s
}

func (s *UnAssociateVpdCidrBlockResponseBody) SetMessage(v string) *UnAssociateVpdCidrBlockResponseBody {
	s.Message = &v
	return s
}

func (s *UnAssociateVpdCidrBlockResponseBody) SetRequestId(v string) *UnAssociateVpdCidrBlockResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnAssociateVpdCidrBlockResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UnAssociateVpdCidrBlockResponseBodyContent struct {
	// The ID of the Lingjun CIDR block.
	//
	// example:
	//
	// vpd-ze3na0wf
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s UnAssociateVpdCidrBlockResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s UnAssociateVpdCidrBlockResponseBodyContent) GoString() string {
	return s.String()
}

func (s *UnAssociateVpdCidrBlockResponseBodyContent) GetVpdId() *string {
	return s.VpdId
}

func (s *UnAssociateVpdCidrBlockResponseBodyContent) SetVpdId(v string) *UnAssociateVpdCidrBlockResponseBodyContent {
	s.VpdId = &v
	return s
}

func (s *UnAssociateVpdCidrBlockResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
