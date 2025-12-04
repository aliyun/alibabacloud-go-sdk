// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDestinationCidrBlockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetDestinationCidrBlockResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *GetDestinationCidrBlockResponseBody
	GetCode() *int32
	SetContent(v *GetDestinationCidrBlockResponseBodyContent) *GetDestinationCidrBlockResponseBody
	GetContent() *GetDestinationCidrBlockResponseBodyContent
	SetMessage(v string) *GetDestinationCidrBlockResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDestinationCidrBlockResponseBody
	GetRequestId() *string
}

type GetDestinationCidrBlockResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response content
	Content *GetDestinationCidrBlockResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// Error message. (Indicates the reason for the anomaly when the instance status is abnormal.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of this request
	//
	// example:
	//
	// D349EE86-AF3F-5F6C-87E2-2A08D3618350
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDestinationCidrBlockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDestinationCidrBlockResponseBody) GoString() string {
	return s.String()
}

func (s *GetDestinationCidrBlockResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetDestinationCidrBlockResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetDestinationCidrBlockResponseBody) GetContent() *GetDestinationCidrBlockResponseBodyContent {
	return s.Content
}

func (s *GetDestinationCidrBlockResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDestinationCidrBlockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDestinationCidrBlockResponseBody) SetAccessDeniedDetail(v string) *GetDestinationCidrBlockResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetDestinationCidrBlockResponseBody) SetCode(v int32) *GetDestinationCidrBlockResponseBody {
	s.Code = &v
	return s
}

func (s *GetDestinationCidrBlockResponseBody) SetContent(v *GetDestinationCidrBlockResponseBodyContent) *GetDestinationCidrBlockResponseBody {
	s.Content = v
	return s
}

func (s *GetDestinationCidrBlockResponseBody) SetMessage(v string) *GetDestinationCidrBlockResponseBody {
	s.Message = &v
	return s
}

func (s *GetDestinationCidrBlockResponseBody) SetRequestId(v string) *GetDestinationCidrBlockResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDestinationCidrBlockResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDestinationCidrBlockResponseBodyContent struct {
	// List of destination CIDR block information for the current network instance
	DestinationCidrBlock []*string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty" type:"Repeated"`
}

func (s GetDestinationCidrBlockResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s GetDestinationCidrBlockResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetDestinationCidrBlockResponseBodyContent) GetDestinationCidrBlock() []*string {
	return s.DestinationCidrBlock
}

func (s *GetDestinationCidrBlockResponseBodyContent) SetDestinationCidrBlock(v []*string) *GetDestinationCidrBlockResponseBodyContent {
	s.DestinationCidrBlock = v
	return s
}

func (s *GetDestinationCidrBlockResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
