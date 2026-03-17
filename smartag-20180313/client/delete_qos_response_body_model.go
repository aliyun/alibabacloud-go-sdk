// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteQosResponseBody
	GetRequestId() *string
}

type DeleteQosResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F77B3F0E-CAA2-41CF-A752-4F2893C5F7F0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteQosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteQosResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteQosResponseBody) SetRequestId(v string) *DeleteQosResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQosResponseBody) Validate() error {
	return dara.Validate(s)
}
