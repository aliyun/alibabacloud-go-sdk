// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateVpdCidrBlockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AssociateVpdCidrBlockResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *AssociateVpdCidrBlockResponseBody
	GetCode() *int32
	SetContent(v *AssociateVpdCidrBlockResponseBodyContent) *AssociateVpdCidrBlockResponseBody
	GetContent() *AssociateVpdCidrBlockResponseBodyContent
	SetMessage(v string) *AssociateVpdCidrBlockResponseBody
	GetMessage() *string
	SetRequestId(v string) *AssociateVpdCidrBlockResponseBody
	GetRequestId() *string
}

type AssociateVpdCidrBlockResponseBody struct {
	// The detailed reason why the access was denied.
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
	Content *AssociateVpdCidrBlockResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
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
	// 9C50C9CD-E799-54DA-BA7A-1FAF3DF80857
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateVpdCidrBlockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateVpdCidrBlockResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateVpdCidrBlockResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AssociateVpdCidrBlockResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AssociateVpdCidrBlockResponseBody) GetContent() *AssociateVpdCidrBlockResponseBodyContent {
	return s.Content
}

func (s *AssociateVpdCidrBlockResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AssociateVpdCidrBlockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateVpdCidrBlockResponseBody) SetAccessDeniedDetail(v string) *AssociateVpdCidrBlockResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AssociateVpdCidrBlockResponseBody) SetCode(v int32) *AssociateVpdCidrBlockResponseBody {
	s.Code = &v
	return s
}

func (s *AssociateVpdCidrBlockResponseBody) SetContent(v *AssociateVpdCidrBlockResponseBodyContent) *AssociateVpdCidrBlockResponseBody {
	s.Content = v
	return s
}

func (s *AssociateVpdCidrBlockResponseBody) SetMessage(v string) *AssociateVpdCidrBlockResponseBody {
	s.Message = &v
	return s
}

func (s *AssociateVpdCidrBlockResponseBody) SetRequestId(v string) *AssociateVpdCidrBlockResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateVpdCidrBlockResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AssociateVpdCidrBlockResponseBodyContent struct {
	// The ID of the Lingjun CIDR block.
	//
	// example:
	//
	// vpd-eoiy88ju
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s AssociateVpdCidrBlockResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s AssociateVpdCidrBlockResponseBodyContent) GoString() string {
	return s.String()
}

func (s *AssociateVpdCidrBlockResponseBodyContent) GetVpdId() *string {
	return s.VpdId
}

func (s *AssociateVpdCidrBlockResponseBodyContent) SetVpdId(v string) *AssociateVpdCidrBlockResponseBodyContent {
	s.VpdId = &v
	return s
}

func (s *AssociateVpdCidrBlockResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
