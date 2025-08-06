// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPlayerConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPlayerConfig(v string) *GetPlayerConfigResponseBody
	GetPlayerConfig() *string
	SetRequestId(v string) *GetPlayerConfigResponseBody
	GetRequestId() *string
}

type GetPlayerConfigResponseBody struct {
	PlayerConfig *string `json:"PlayerConfig,omitempty" xml:"PlayerConfig,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPlayerConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPlayerConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetPlayerConfigResponseBody) GetPlayerConfig() *string {
	return s.PlayerConfig
}

func (s *GetPlayerConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPlayerConfigResponseBody) SetPlayerConfig(v string) *GetPlayerConfigResponseBody {
	s.PlayerConfig = &v
	return s
}

func (s *GetPlayerConfigResponseBody) SetRequestId(v string) *GetPlayerConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPlayerConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
