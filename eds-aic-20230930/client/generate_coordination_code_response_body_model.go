// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCoordinationCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCoordinatorCode(v string) *GenerateCoordinationCodeResponseBody
	GetCoordinatorCode() *string
	SetRequestId(v string) *GenerateCoordinationCodeResponseBody
	GetRequestId() *string
}

type GenerateCoordinationCodeResponseBody struct {
	// The collaboration code.
	//
	// example:
	//
	// CSHGDK
	CoordinatorCode *string `json:"CoordinatorCode,omitempty" xml:"CoordinatorCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1A923337-44D9-5CAD-9A53-95084BD4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateCoordinationCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateCoordinationCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateCoordinationCodeResponseBody) GetCoordinatorCode() *string {
	return s.CoordinatorCode
}

func (s *GenerateCoordinationCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateCoordinationCodeResponseBody) SetCoordinatorCode(v string) *GenerateCoordinationCodeResponseBody {
	s.CoordinatorCode = &v
	return s
}

func (s *GenerateCoordinationCodeResponseBody) SetRequestId(v string) *GenerateCoordinationCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateCoordinationCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
