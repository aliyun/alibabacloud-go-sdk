// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCasterEpisodeGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCasterEpisodeGroupResponseBody
	GetRequestId() *string
}

type DeleteCasterEpisodeGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCasterEpisodeGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCasterEpisodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCasterEpisodeGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCasterEpisodeGroupResponseBody) SetRequestId(v string) *DeleteCasterEpisodeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCasterEpisodeGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
