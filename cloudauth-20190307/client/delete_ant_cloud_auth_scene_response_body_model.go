// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAntCloudAuthSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAntCloudAuthSceneResponseBody
	GetRequestId() *string
}

type DeleteAntCloudAuthSceneResponseBody struct {
	// ID of this request.
	//
	// example:
	//
	// 80D1ACD4-1C7D-53CB-8C54-3758FF03C762
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAntCloudAuthSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAntCloudAuthSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAntCloudAuthSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAntCloudAuthSceneResponseBody) SetRequestId(v string) *DeleteAntCloudAuthSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAntCloudAuthSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
