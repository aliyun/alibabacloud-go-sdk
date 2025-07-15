// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCasterSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopCasterSceneResponseBody
	GetRequestId() *string
}

type StopCasterSceneResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CF60DB6A-7FD6-426E-9288-122CC1A52FA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopCasterSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopCasterSceneResponseBody) GoString() string {
	return s.String()
}

func (s *StopCasterSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopCasterSceneResponseBody) SetRequestId(v string) *StopCasterSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopCasterSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
