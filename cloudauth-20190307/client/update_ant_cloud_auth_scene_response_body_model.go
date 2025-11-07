// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAntCloudAuthSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAntCloudAuthSceneResponseBody
	GetRequestId() *string
}

type UpdateAntCloudAuthSceneResponseBody struct {
	// ID of this request.
	//
	// example:
	//
	// A07D5838-2BE0-5D2E-B864-D05ADAA5EDE3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAntCloudAuthSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAntCloudAuthSceneResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAntCloudAuthSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAntCloudAuthSceneResponseBody) SetRequestId(v string) *UpdateAntCloudAuthSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAntCloudAuthSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
