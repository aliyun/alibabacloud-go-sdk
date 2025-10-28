// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEcuRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEcuId(v string) *DeleteEcuRequest
	GetEcuId() *string
}

type DeleteEcuRequest struct {
	// The unique ID of the ECU to be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5c0b8c82-4ba9-****-****-130a34ffa534
	EcuId *string `json:"EcuId,omitempty" xml:"EcuId,omitempty"`
}

func (s DeleteEcuRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEcuRequest) GoString() string {
	return s.String()
}

func (s *DeleteEcuRequest) GetEcuId() *string {
	return s.EcuId
}

func (s *DeleteEcuRequest) SetEcuId(v string) *DeleteEcuRequest {
	s.EcuId = &v
	return s
}

func (s *DeleteEcuRequest) Validate() error {
	return dara.Validate(s)
}
