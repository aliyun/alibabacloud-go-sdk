// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExternalDataServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteExternalDataServiceResponseBody
	GetRequestId() *string
}

type DeleteExternalDataServiceResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteExternalDataServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteExternalDataServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExternalDataServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteExternalDataServiceResponseBody) SetRequestId(v string) *DeleteExternalDataServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExternalDataServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
