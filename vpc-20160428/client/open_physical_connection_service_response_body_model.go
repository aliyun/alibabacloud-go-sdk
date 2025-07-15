// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenPhysicalConnectionServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OpenPhysicalConnectionServiceResponseBody
	GetRequestId() *string
}

type OpenPhysicalConnectionServiceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenPhysicalConnectionServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenPhysicalConnectionServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenPhysicalConnectionServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenPhysicalConnectionServiceResponseBody) SetRequestId(v string) *OpenPhysicalConnectionServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenPhysicalConnectionServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
