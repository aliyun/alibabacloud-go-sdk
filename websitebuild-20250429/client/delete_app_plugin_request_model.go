// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppPluginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPluginId(v string) *DeleteAppPluginRequest
	GetPluginId() *string
}

type DeleteAppPluginRequest struct {
	// The gateway plug-in ID.
	//
	// example:
	//
	// aliplayer-react
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
}

func (s DeleteAppPluginRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppPluginRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppPluginRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *DeleteAppPluginRequest) SetPluginId(v string) *DeleteAppPluginRequest {
	s.PluginId = &v
	return s
}

func (s *DeleteAppPluginRequest) Validate() error {
	return dara.Validate(s)
}
