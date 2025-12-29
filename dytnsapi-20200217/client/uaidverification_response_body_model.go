// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUAIDVerificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UAIDVerificationResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UAIDVerificationResponseBody
	GetCode() *string
	SetData(v *UAIDVerificationResponseBodyData) *UAIDVerificationResponseBody
	GetData() *UAIDVerificationResponseBodyData
	SetMessage(v string) *UAIDVerificationResponseBody
	GetMessage() *string
	SetRequestId(v string) *UAIDVerificationResponseBody
	GetRequestId() *string
}

type UAIDVerificationResponseBody struct {
	// example:
	//
	// -
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UAIDVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 68A40250-50CD-034C-B728-0BD******177
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UAIDVerificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UAIDVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *UAIDVerificationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UAIDVerificationResponseBody) GetCode() *string {
	return s.Code
}

func (s *UAIDVerificationResponseBody) GetData() *UAIDVerificationResponseBodyData {
	return s.Data
}

func (s *UAIDVerificationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UAIDVerificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UAIDVerificationResponseBody) SetAccessDeniedDetail(v string) *UAIDVerificationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UAIDVerificationResponseBody) SetCode(v string) *UAIDVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *UAIDVerificationResponseBody) SetData(v *UAIDVerificationResponseBodyData) *UAIDVerificationResponseBody {
	s.Data = v
	return s
}

func (s *UAIDVerificationResponseBody) SetMessage(v string) *UAIDVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *UAIDVerificationResponseBody) SetRequestId(v string) *UAIDVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UAIDVerificationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UAIDVerificationResponseBodyData struct {
	// example:
	//
	// B1E0C1********9F757AF52A035
	Uaid *string `json:"Uaid,omitempty" xml:"Uaid,omitempty"`
}

func (s UAIDVerificationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UAIDVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *UAIDVerificationResponseBodyData) GetUaid() *string {
	return s.Uaid
}

func (s *UAIDVerificationResponseBodyData) SetUaid(v string) *UAIDVerificationResponseBodyData {
	s.Uaid = &v
	return s
}

func (s *UAIDVerificationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
