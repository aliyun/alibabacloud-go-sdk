// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSceneResponseBody
	GetRequestId() *string
	SetSceneId(v string) *CreateSceneResponseBody
	GetSceneId() *string
}

type CreateSceneResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// FCF741D8-9C30-578E-807F-B935487DB34A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s CreateSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSceneResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSceneResponseBody) GetSceneId() *string {
	return s.SceneId
}

func (s *CreateSceneResponseBody) SetRequestId(v string) *CreateSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSceneResponseBody) SetSceneId(v string) *CreateSceneResponseBody {
	s.SceneId = &v
	return s
}

func (s *CreateSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
