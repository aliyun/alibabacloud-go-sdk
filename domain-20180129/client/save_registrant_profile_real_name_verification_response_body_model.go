// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveRegistrantProfileRealNameVerificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegistrantProfileId(v int64) *SaveRegistrantProfileRealNameVerificationResponseBody
	GetRegistrantProfileId() *int64
	SetRequestId(v string) *SaveRegistrantProfileRealNameVerificationResponseBody
	GetRequestId() *string
}

type SaveRegistrantProfileRealNameVerificationResponseBody struct {
	// example:
	//
	// 1234567
	RegistrantProfileId *int64 `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	// example:
	//
	// 4D73432C-7600-****-ACBB-C3B5CA145D32
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveRegistrantProfileRealNameVerificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveRegistrantProfileRealNameVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *SaveRegistrantProfileRealNameVerificationResponseBody) GetRegistrantProfileId() *int64 {
	return s.RegistrantProfileId
}

func (s *SaveRegistrantProfileRealNameVerificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveRegistrantProfileRealNameVerificationResponseBody) SetRegistrantProfileId(v int64) *SaveRegistrantProfileRealNameVerificationResponseBody {
	s.RegistrantProfileId = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationResponseBody) SetRequestId(v string) *SaveRegistrantProfileRealNameVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationResponseBody) Validate() error {
	return dara.Validate(s)
}
