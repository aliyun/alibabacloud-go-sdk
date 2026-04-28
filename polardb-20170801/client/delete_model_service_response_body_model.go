// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteModelServiceResponseBody
	GetRequestId() *string
}

type DeleteModelServiceResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// CDB3258F-B5DE-43C4-8935-CBA0CA******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteModelServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModelServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteModelServiceResponseBody) SetRequestId(v string) *DeleteModelServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteModelServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
