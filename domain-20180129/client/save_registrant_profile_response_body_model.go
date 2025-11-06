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
	// 3600000
	RegistrantProfileId *int64 `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	// example:
	//
	// D09B153B-294D-42F1-BB61-F1C72136DFD3
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
