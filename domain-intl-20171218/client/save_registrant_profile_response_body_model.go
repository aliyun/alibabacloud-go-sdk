// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveRegistrantProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegistrantProfileId(v int64) *SaveRegistrantProfileResponseBody
	GetRegistrantProfileId() *int64
	SetRequestId(v string) *SaveRegistrantProfileResponseBody
	GetRequestId() *string
}

type SaveRegistrantProfileResponseBody struct {
	// example:
	//
	// 12380891
	RegistrantProfileId *int64 `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	// example:
	//
	// A9C35C47-3366-482E-B872-8C9EA4733FE9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveRegistrantProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveRegistrantProfileResponseBody) GoString() string {
	return s.String()
}

func (s *SaveRegistrantProfileResponseBody) GetRegistrantProfileId() *int64 {
	return s.RegistrantProfileId
}

func (s *SaveRegistrantProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveRegistrantProfileResponseBody) SetRegistrantProfileId(v int64) *SaveRegistrantProfileResponseBody {
	s.RegistrantProfileId = &v
	return s
}

func (s *SaveRegistrantProfileResponseBody) SetRequestId(v string) *SaveRegistrantProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveRegistrantProfileResponseBody) Validate() error {
	return dara.Validate(s)
}
