// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateVpdResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *CreateVpdResponseBody
	GetCode() *int32
	SetContent(v *CreateVpdResponseBodyContent) *CreateVpdResponseBody
	GetContent() *CreateVpdResponseBodyContent
	SetMessage(v string) *CreateVpdResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateVpdResponseBody
	GetRequestId() *string
}

type CreateVpdResponseBody struct {
	// The details about the failed permission verification.
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
	Content *CreateVpdResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DBAD15D6-3F47-5B36-8A92-57C2919D13D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVpdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVpdResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpdResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateVpdResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateVpdResponseBody) GetContent() *CreateVpdResponseBodyContent {
	return s.Content
}

func (s *CreateVpdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateVpdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVpdResponseBody) SetAccessDeniedDetail(v string) *CreateVpdResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateVpdResponseBody) SetCode(v int32) *CreateVpdResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVpdResponseBody) SetContent(v *CreateVpdResponseBodyContent) *CreateVpdResponseBody {
	s.Content = v
	return s
}

func (s *CreateVpdResponseBody) SetMessage(v string) *CreateVpdResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVpdResponseBody) SetRequestId(v string) *CreateVpdResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpdResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateVpdResponseBodyContent struct {
	// Lingjun subnet ID list
	SubnetIds []*string `json:"SubnetIds,omitempty" xml:"SubnetIds,omitempty" type:"Repeated"`
	// Lingjun CIDR block instance ID
	//
	// example:
	//
	// vpd-eoiy88ju
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s CreateVpdResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s CreateVpdResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CreateVpdResponseBodyContent) GetSubnetIds() []*string {
	return s.SubnetIds
}

func (s *CreateVpdResponseBodyContent) GetVpdId() *string {
	return s.VpdId
}

func (s *CreateVpdResponseBodyContent) SetSubnetIds(v []*string) *CreateVpdResponseBodyContent {
	s.SubnetIds = v
	return s
}

func (s *CreateVpdResponseBodyContent) SetVpdId(v string) *CreateVpdResponseBodyContent {
	s.VpdId = &v
	return s
}

func (s *CreateVpdResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
