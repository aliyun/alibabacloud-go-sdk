// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCasterSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartCasterSceneResponseBody
	GetRequestId() *string
	SetStreamUrl(v string) *StartCasterSceneResponseBody
	GetStreamUrl() *string
}

type StartCasterSceneResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CF60DB6A-7FD6-426E-9288-122CC1A52FA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The streaming URL of the scene. It is used for playback, but not for stream relay.
	//
	// example:
	//
	// http://live/caster/example.org
	StreamUrl *string `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty"`
}

func (s StartCasterSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartCasterSceneResponseBody) GoString() string {
	return s.String()
}

func (s *StartCasterSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartCasterSceneResponseBody) GetStreamUrl() *string {
	return s.StreamUrl
}

func (s *StartCasterSceneResponseBody) SetRequestId(v string) *StartCasterSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartCasterSceneResponseBody) SetStreamUrl(v string) *StartCasterSceneResponseBody {
	s.StreamUrl = &v
	return s
}

func (s *StartCasterSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
