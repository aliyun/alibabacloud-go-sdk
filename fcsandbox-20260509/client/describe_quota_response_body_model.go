// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeQuotaResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeQuotaResponseBody
	GetMessage() *string
	SetQuota(v *Quota) *DescribeQuotaResponseBody
	GetQuota() *Quota
	SetRequestId(v string) *DescribeQuotaResponseBody
	GetRequestId() *string
}

type DescribeQuotaResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The Quota configuration information.
	Quota *Quota `json:"quota,omitempty" xml:"quota,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7ADFF8D8-D4BA-5F79-AD49-DDABFEA59B6C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DescribeQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQuotaResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeQuotaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeQuotaResponseBody) GetQuota() *Quota {
	return s.Quota
}

func (s *DescribeQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeQuotaResponseBody) SetCode(v string) *DescribeQuotaResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeQuotaResponseBody) SetMessage(v string) *DescribeQuotaResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeQuotaResponseBody) SetQuota(v *Quota) *DescribeQuotaResponseBody {
	s.Quota = v
	return s
}

func (s *DescribeQuotaResponseBody) SetRequestId(v string) *DescribeQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeQuotaResponseBody) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}
