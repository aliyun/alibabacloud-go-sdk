// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebPathRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ModifyWebPathRequest
	GetConfig() *string
	SetTarget(v string) *ModifyWebPathRequest
	GetTarget() *string
	SetType(v string) *ModifyWebPathRequest
	GetType() *string
}

type ModifyWebPathRequest struct {
	// The configuration of the web directory. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **webPathType**: the type of the web directory
	//
	// 	- **webPath**: the web directory
	//
	// example:
	//
	// {
	//
	//       "webPathType": "customize",
	//
	//       "webPath": "/root/www****"
	//
	// }
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The protected asset to which the web directory belongs. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **target**: the protected asset.
	//
	// 	- **targetType**: the type of the asset. Set the value to uuid.
	//
	// 	- **flag**: the type of the operation.
	//
	// example:
	//
	// [{"target":"0186127a-d33e-4d0c-94fb-8f25f87bc69f","targetType":"uuid","flag":"add"}]
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The type of the configuration item. Set the value to **web_path**.
	//
	// example:
	//
	// web_path
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyWebPathRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebPathRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebPathRequest) GetConfig() *string {
	return s.Config
}

func (s *ModifyWebPathRequest) GetTarget() *string {
	return s.Target
}

func (s *ModifyWebPathRequest) GetType() *string {
	return s.Type
}

func (s *ModifyWebPathRequest) SetConfig(v string) *ModifyWebPathRequest {
	s.Config = &v
	return s
}

func (s *ModifyWebPathRequest) SetTarget(v string) *ModifyWebPathRequest {
	s.Target = &v
	return s
}

func (s *ModifyWebPathRequest) SetType(v string) *ModifyWebPathRequest {
	s.Type = &v
	return s
}

func (s *ModifyWebPathRequest) Validate() error {
	return dara.Validate(s)
}
