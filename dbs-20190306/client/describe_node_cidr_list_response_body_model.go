// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodeCidrListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DescribeNodeCidrListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeNodeCidrListResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeNodeCidrListResponseBody
	GetHttpStatusCode() *int32
	SetInternetIPs(v *DescribeNodeCidrListResponseBodyInternetIPs) *DescribeNodeCidrListResponseBody
	GetInternetIPs() *DescribeNodeCidrListResponseBodyInternetIPs
	SetIntranetIPs(v *DescribeNodeCidrListResponseBodyIntranetIPs) *DescribeNodeCidrListResponseBody
	GetIntranetIPs() *DescribeNodeCidrListResponseBodyIntranetIPs
	SetRequestId(v string) *DescribeNodeCidrListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeNodeCidrListResponseBody
	GetSuccess() *bool
}

type DescribeNodeCidrListResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Param.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The public CIDR blocks.
	InternetIPs *DescribeNodeCidrListResponseBodyInternetIPs `json:"InternetIPs,omitempty" xml:"InternetIPs,omitempty" type:"Struct"`
	// The internal CIDR blocks.
	IntranetIPs *DescribeNodeCidrListResponseBodyIntranetIPs `json:"IntranetIPs,omitempty" xml:"IntranetIPs,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 5B352E69-E7B1-4EA7-BB8E-29FFE969C791
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeNodeCidrListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeCidrListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNodeCidrListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeNodeCidrListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeNodeCidrListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeNodeCidrListResponseBody) GetInternetIPs() *DescribeNodeCidrListResponseBodyInternetIPs {
	return s.InternetIPs
}

func (s *DescribeNodeCidrListResponseBody) GetIntranetIPs() *DescribeNodeCidrListResponseBodyIntranetIPs {
	return s.IntranetIPs
}

func (s *DescribeNodeCidrListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNodeCidrListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeNodeCidrListResponseBody) SetErrCode(v string) *DescribeNodeCidrListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeNodeCidrListResponseBody) SetErrMessage(v string) *DescribeNodeCidrListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeNodeCidrListResponseBody) SetHttpStatusCode(v int32) *DescribeNodeCidrListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeNodeCidrListResponseBody) SetInternetIPs(v *DescribeNodeCidrListResponseBodyInternetIPs) *DescribeNodeCidrListResponseBody {
	s.InternetIPs = v
	return s
}

func (s *DescribeNodeCidrListResponseBody) SetIntranetIPs(v *DescribeNodeCidrListResponseBodyIntranetIPs) *DescribeNodeCidrListResponseBody {
	s.IntranetIPs = v
	return s
}

func (s *DescribeNodeCidrListResponseBody) SetRequestId(v string) *DescribeNodeCidrListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNodeCidrListResponseBody) SetSuccess(v bool) *DescribeNodeCidrListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeNodeCidrListResponseBody) Validate() error {
	if s.InternetIPs != nil {
		if err := s.InternetIPs.Validate(); err != nil {
			return err
		}
	}
	if s.IntranetIPs != nil {
		if err := s.IntranetIPs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNodeCidrListResponseBodyInternetIPs struct {
	InternetIP []*string `json:"InternetIP,omitempty" xml:"InternetIP,omitempty" type:"Repeated"`
}

func (s DescribeNodeCidrListResponseBodyInternetIPs) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeCidrListResponseBodyInternetIPs) GoString() string {
	return s.String()
}

func (s *DescribeNodeCidrListResponseBodyInternetIPs) GetInternetIP() []*string {
	return s.InternetIP
}

func (s *DescribeNodeCidrListResponseBodyInternetIPs) SetInternetIP(v []*string) *DescribeNodeCidrListResponseBodyInternetIPs {
	s.InternetIP = v
	return s
}

func (s *DescribeNodeCidrListResponseBodyInternetIPs) Validate() error {
	return dara.Validate(s)
}

type DescribeNodeCidrListResponseBodyIntranetIPs struct {
	IntranetIP []*string `json:"IntranetIP,omitempty" xml:"IntranetIP,omitempty" type:"Repeated"`
}

func (s DescribeNodeCidrListResponseBodyIntranetIPs) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeCidrListResponseBodyIntranetIPs) GoString() string {
	return s.String()
}

func (s *DescribeNodeCidrListResponseBodyIntranetIPs) GetIntranetIP() []*string {
	return s.IntranetIP
}

func (s *DescribeNodeCidrListResponseBodyIntranetIPs) SetIntranetIP(v []*string) *DescribeNodeCidrListResponseBodyIntranetIPs {
	s.IntranetIP = v
	return s
}

func (s *DescribeNodeCidrListResponseBodyIntranetIPs) Validate() error {
	return dara.Validate(s)
}
