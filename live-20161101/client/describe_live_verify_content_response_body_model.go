// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveVerifyContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *DescribeLiveVerifyContentResponseBody
	GetContent() *string
	SetRequestId(v string) *DescribeLiveVerifyContentResponseBody
	GetRequestId() *string
}

type DescribeLiveVerifyContentResponseBody struct {
	// The verification content.
	//
	// example:
	//
	// verify_dffeb6610035dcb77b413******
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveVerifyContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveVerifyContentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveVerifyContentResponseBody) GetContent() *string {
	return s.Content
}

func (s *DescribeLiveVerifyContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveVerifyContentResponseBody) SetContent(v string) *DescribeLiveVerifyContentResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeLiveVerifyContentResponseBody) SetRequestId(v string) *DescribeLiveVerifyContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveVerifyContentResponseBody) Validate() error {
	return dara.Validate(s)
}
