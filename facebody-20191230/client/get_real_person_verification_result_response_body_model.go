// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRealPersonVerificationResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetRealPersonVerificationResultResponseBodyData) *GetRealPersonVerificationResultResponseBody
	GetData() *GetRealPersonVerificationResultResponseBodyData
	SetRequestId(v string) *GetRealPersonVerificationResultResponseBody
	GetRequestId() *string
}

type GetRealPersonVerificationResultResponseBody struct {
	Data *GetRealPersonVerificationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 10FC953D-7C0C-4915-8949-34E3246E5B79
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRealPersonVerificationResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRealPersonVerificationResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetRealPersonVerificationResultResponseBody) GetData() *GetRealPersonVerificationResultResponseBodyData {
	return s.Data
}

func (s *GetRealPersonVerificationResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRealPersonVerificationResultResponseBody) SetData(v *GetRealPersonVerificationResultResponseBodyData) *GetRealPersonVerificationResultResponseBody {
	s.Data = v
	return s
}

func (s *GetRealPersonVerificationResultResponseBody) SetRequestId(v string) *GetRealPersonVerificationResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRealPersonVerificationResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRealPersonVerificationResultResponseBodyData struct {
	// example:
	//
	// true
	Passed *bool `json:"Passed,omitempty" xml:"Passed,omitempty"`
}

func (s GetRealPersonVerificationResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRealPersonVerificationResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRealPersonVerificationResultResponseBodyData) GetPassed() *bool {
	return s.Passed
}

func (s *GetRealPersonVerificationResultResponseBodyData) SetPassed(v bool) *GetRealPersonVerificationResultResponseBodyData {
	s.Passed = &v
	return s
}

func (s *GetRealPersonVerificationResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
