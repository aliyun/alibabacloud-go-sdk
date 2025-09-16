// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdvanceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *GetAdvanceConfigRequest
	GetType() *string
}

type GetAdvanceConfigRequest struct {
	// 	- The type of the advanced configuration. Valid values: -ONLINE: online configuration
	//
	// 	- \\-ONLINE_CAVA: online Cava configuration
	//
	// 	- \\-ONLINE_PLUGIN: online plug-in configuration
	//
	// 	- \\-ONLINE_QUERY: query configuration
	//
	// 	- \\-OFFLINE_DICT: offline dictionary configuration
	//
	// 	- \\-OFFLINE_TABLE: offline table configuration
	//
	// 	- \\-OFFLINE_COMMON: offline configuration
	//
	// 	- \\-OFFLINE_PLUGIN: offline plug-in configuration
	//
	// 	- \\-OFFLINE_INDEX: index configuration
	//
	// example:
	//
	// ONLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetAdvanceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAdvanceConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigRequest) GetType() *string {
	return s.Type
}

func (s *GetAdvanceConfigRequest) SetType(v string) *GetAdvanceConfigRequest {
	s.Type = &v
	return s
}

func (s *GetAdvanceConfigRequest) Validate() error {
	return dara.Validate(s)
}
