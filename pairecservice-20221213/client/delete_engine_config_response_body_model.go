// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEngineConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEngineConfigResponseBody
	GetRequestId() *string
}

type DeleteEngineConfigResponseBody struct {
	// example:
	//
	// F7AC05FF-EDE7-5C2B-B9AE-33D6DF4178BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEngineConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEngineConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEngineConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEngineConfigResponseBody) SetRequestId(v string) *DeleteEngineConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEngineConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
