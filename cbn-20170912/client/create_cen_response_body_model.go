// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *CreateCenResponseBody
	GetCenId() *string
	SetRequestId(v string) *CreateCenResponseBody
	GetRequestId() *string
}

type CreateCenResponseBody struct {
	// The CEN instance ID.
	//
	// example:
	//
	// cen-dc4vwznpwbobrl****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0C2EE7A8-74D4-4081-8236-CEBDE3BBCF50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCenResponseBody) GetCenId() *string {
	return s.CenId
}

func (s *CreateCenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCenResponseBody) SetCenId(v string) *CreateCenResponseBody {
	s.CenId = &v
	return s
}

func (s *CreateCenResponseBody) SetRequestId(v string) *CreateCenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCenResponseBody) Validate() error {
	return dara.Validate(s)
}
