// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKyuubiServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateKyuubiServiceResponseBodyData) *CreateKyuubiServiceResponseBody
	GetData() *CreateKyuubiServiceResponseBodyData
	SetRequestId(v string) *CreateKyuubiServiceResponseBody
	GetRequestId() *string
}

type CreateKyuubiServiceResponseBody struct {
	Data *CreateKyuubiServiceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateKyuubiServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKyuubiServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKyuubiServiceResponseBody) GetData() *CreateKyuubiServiceResponseBodyData {
	return s.Data
}

func (s *CreateKyuubiServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKyuubiServiceResponseBody) SetData(v *CreateKyuubiServiceResponseBodyData) *CreateKyuubiServiceResponseBody {
	s.Data = v
	return s
}

func (s *CreateKyuubiServiceResponseBody) SetRequestId(v string) *CreateKyuubiServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKyuubiServiceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateKyuubiServiceResponseBodyData struct {
	// Kyuubi Service ID。
	//
	// example:
	//
	// kb-f99935b350fb4****7ef700b8b4197a3
	KyuubiServiceId *string `json:"kyuubiServiceId,omitempty" xml:"kyuubiServiceId,omitempty"`
}

func (s CreateKyuubiServiceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateKyuubiServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateKyuubiServiceResponseBodyData) GetKyuubiServiceId() *string {
	return s.KyuubiServiceId
}

func (s *CreateKyuubiServiceResponseBodyData) SetKyuubiServiceId(v string) *CreateKyuubiServiceResponseBodyData {
	s.KyuubiServiceId = &v
	return s
}

func (s *CreateKyuubiServiceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
