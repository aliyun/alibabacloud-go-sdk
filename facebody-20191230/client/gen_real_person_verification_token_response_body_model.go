// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenRealPersonVerificationTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GenRealPersonVerificationTokenResponseBodyData) *GenRealPersonVerificationTokenResponseBody
	GetData() *GenRealPersonVerificationTokenResponseBodyData
	SetRequestId(v string) *GenRealPersonVerificationTokenResponseBody
	GetRequestId() *string
}

type GenRealPersonVerificationTokenResponseBody struct {
	Data *GenRealPersonVerificationTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// A31006F2-22E7-4538-93BB-DE6B563643EE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenRealPersonVerificationTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenRealPersonVerificationTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GenRealPersonVerificationTokenResponseBody) GetData() *GenRealPersonVerificationTokenResponseBodyData {
	return s.Data
}

func (s *GenRealPersonVerificationTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenRealPersonVerificationTokenResponseBody) SetData(v *GenRealPersonVerificationTokenResponseBodyData) *GenRealPersonVerificationTokenResponseBody {
	s.Data = v
	return s
}

func (s *GenRealPersonVerificationTokenResponseBody) SetRequestId(v string) *GenRealPersonVerificationTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenRealPersonVerificationTokenResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenRealPersonVerificationTokenResponseBodyData struct {
	// example:
	//
	// 9fca3791c158a479ead9f2ba65ab3XXX
	VerificationToken *string `json:"VerificationToken,omitempty" xml:"VerificationToken,omitempty"`
}

func (s GenRealPersonVerificationTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenRealPersonVerificationTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenRealPersonVerificationTokenResponseBodyData) GetVerificationToken() *string {
	return s.VerificationToken
}

func (s *GenRealPersonVerificationTokenResponseBodyData) SetVerificationToken(v string) *GenRealPersonVerificationTokenResponseBodyData {
	s.VerificationToken = &v
	return s
}

func (s *GenRealPersonVerificationTokenResponseBodyData) Validate() error {
	return dara.Validate(s)
}
