// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAntCloudAuthSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAntCloudAuthSceneResponseBody
	GetRequestId() *string
	SetSceneId(v int64) *CreateAntCloudAuthSceneResponseBody
	GetSceneId() *int64
}

type CreateAntCloudAuthSceneResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 01D3BDC6-64C0-58E2-8760-3F1B56AAE299
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// 1000015112
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s CreateAntCloudAuthSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAntCloudAuthSceneResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAntCloudAuthSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAntCloudAuthSceneResponseBody) GetSceneId() *int64 {
	return s.SceneId
}

func (s *CreateAntCloudAuthSceneResponseBody) SetRequestId(v string) *CreateAntCloudAuthSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAntCloudAuthSceneResponseBody) SetSceneId(v int64) *CreateAntCloudAuthSceneResponseBody {
	s.SceneId = &v
	return s
}

func (s *CreateAntCloudAuthSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
