// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCasterSceneAudioResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCasterSceneAudioResponseBody
	GetRequestId() *string
}

type UpdateCasterSceneAudioResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCasterSceneAudioResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCasterSceneAudioResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCasterSceneAudioResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCasterSceneAudioResponseBody) SetRequestId(v string) *UpdateCasterSceneAudioResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCasterSceneAudioResponseBody) Validate() error {
	return dara.Validate(s)
}
