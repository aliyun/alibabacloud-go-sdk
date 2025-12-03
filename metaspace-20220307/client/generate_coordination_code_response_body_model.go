// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCoordinationCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCoordinationCode(v string) *GenerateCoordinationCodeResponseBody
	GetCoordinationCode() *string
	SetRequestId(v string) *GenerateCoordinationCodeResponseBody
	GetRequestId() *string
}

type GenerateCoordinationCodeResponseBody struct {
	// example:
	//
	// PA3MU***
	CoordinationCode *string `json:"CoordinationCode,omitempty" xml:"CoordinationCode,omitempty"`
	// example:
	//
	// AD2D0761-1FE5-549D-B169******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateCoordinationCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateCoordinationCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateCoordinationCodeResponseBody) GetCoordinationCode() *string {
	return s.CoordinationCode
}

func (s *GenerateCoordinationCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateCoordinationCodeResponseBody) SetCoordinationCode(v string) *GenerateCoordinationCodeResponseBody {
	s.CoordinationCode = &v
	return s
}

func (s *GenerateCoordinationCodeResponseBody) SetRequestId(v string) *GenerateCoordinationCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateCoordinationCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
