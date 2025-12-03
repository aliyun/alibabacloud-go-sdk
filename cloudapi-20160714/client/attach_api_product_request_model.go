// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachApiProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiProductId(v string) *AttachApiProductRequest
	GetApiProductId() *string
	SetApis(v []*AttachApiProductRequestApis) *AttachApiProductRequest
	GetApis() []*AttachApiProductRequestApis
	SetSecurityToken(v string) *AttachApiProductRequest
	GetSecurityToken() *string
}

type AttachApiProductRequest struct {
	// The ID of the API product.
	//
	// This parameter is required.
	//
	// example:
	//
	// 117b7a64a8b3f064eaa4a47ac62aac5e
	ApiProductId *string `json:"ApiProductId,omitempty" xml:"ApiProductId,omitempty"`
	// The APIs to be attached.
	//
	// This parameter is required.
	Apis          []*AttachApiProductRequestApis `json:"Apis,omitempty" xml:"Apis,omitempty" type:"Repeated"`
	SecurityToken *string                        `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s AttachApiProductRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachApiProductRequest) GoString() string {
	return s.String()
}

func (s *AttachApiProductRequest) GetApiProductId() *string {
	return s.ApiProductId
}

func (s *AttachApiProductRequest) GetApis() []*AttachApiProductRequestApis {
	return s.Apis
}

func (s *AttachApiProductRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AttachApiProductRequest) SetApiProductId(v string) *AttachApiProductRequest {
	s.ApiProductId = &v
	return s
}

func (s *AttachApiProductRequest) SetApis(v []*AttachApiProductRequestApis) *AttachApiProductRequest {
	s.Apis = v
	return s
}

func (s *AttachApiProductRequest) SetSecurityToken(v string) *AttachApiProductRequest {
	s.SecurityToken = &v
	return s
}

func (s *AttachApiProductRequest) Validate() error {
	if s.Apis != nil {
		for _, item := range s.Apis {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AttachApiProductRequestApis struct {
	// The API ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 551877242a4b4f3a84a56b7c3570e4a7
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The environment. Valid values:
	//
	// 	- **RELEASE**: the production environment
	//
	// 	- **PRE**: the staging environment
	//
	// 	- **TEST**: the test environment
	//
	// This parameter is required.
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s AttachApiProductRequestApis) String() string {
	return dara.Prettify(s)
}

func (s AttachApiProductRequestApis) GoString() string {
	return s.String()
}

func (s *AttachApiProductRequestApis) GetApiId() *string {
	return s.ApiId
}

func (s *AttachApiProductRequestApis) GetStageName() *string {
	return s.StageName
}

func (s *AttachApiProductRequestApis) SetApiId(v string) *AttachApiProductRequestApis {
	s.ApiId = &v
	return s
}

func (s *AttachApiProductRequestApis) SetStageName(v string) *AttachApiProductRequestApis {
	s.StageName = &v
	return s
}

func (s *AttachApiProductRequestApis) Validate() error {
	return dara.Validate(s)
}
