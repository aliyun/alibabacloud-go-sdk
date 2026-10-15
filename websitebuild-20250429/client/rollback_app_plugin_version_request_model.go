// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackAppPluginVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPluginId(v string) *RollbackAppPluginVersionRequest
	GetPluginId() *string
	SetPluginVersion(v string) *RollbackAppPluginVersionRequest
	GetPluginVersion() *string
}

type RollbackAppPluginVersionRequest struct {
	// The gateway plug-in ID.
	//
	// example:
	//
	// 1bae9ceaceea432d91c7069fab0dfc02
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// The target version number.
	//
	// example:
	//
	// v2
	PluginVersion *string `json:"PluginVersion,omitempty" xml:"PluginVersion,omitempty"`
}

func (s RollbackAppPluginVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackAppPluginVersionRequest) GoString() string {
	return s.String()
}

func (s *RollbackAppPluginVersionRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *RollbackAppPluginVersionRequest) GetPluginVersion() *string {
	return s.PluginVersion
}

func (s *RollbackAppPluginVersionRequest) SetPluginId(v string) *RollbackAppPluginVersionRequest {
	s.PluginId = &v
	return s
}

func (s *RollbackAppPluginVersionRequest) SetPluginVersion(v string) *RollbackAppPluginVersionRequest {
	s.PluginVersion = &v
	return s
}

func (s *RollbackAppPluginVersionRequest) Validate() error {
	return dara.Validate(s)
}
