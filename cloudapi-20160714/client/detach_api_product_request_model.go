// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachApiProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiProductId(v string) *DetachApiProductRequest
	GetApiProductId() *string
	SetApis(v []*DetachApiProductRequestApis) *DetachApiProductRequest
	GetApis() []*DetachApiProductRequestApis
	SetSecurityToken(v string) *DetachApiProductRequest
	GetSecurityToken() *string
}

type DetachApiProductRequest struct {
	// The ID of the API product.
	//
	// This parameter is required.
	//
	// example:
	//
	// 117b7a64a8b3f064eaa4a47ac62aac5e
	ApiProductId *string `json:"ApiProductId,omitempty" xml:"ApiProductId,omitempty"`
	// The APIs that you want to detach from the API product.
	//
	// This parameter is required.
	Apis          []*DetachApiProductRequestApis `json:"Apis,omitempty" xml:"Apis,omitempty" type:"Repeated"`
	SecurityToken *string                        `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DetachApiProductRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachApiProductRequest) GoString() string {
	return s.String()
}

func (s *DetachApiProductRequest) GetApiProductId() *string {
	return s.ApiProductId
}

func (s *DetachApiProductRequest) GetApis() []*DetachApiProductRequestApis {
	return s.Apis
}

func (s *DetachApiProductRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DetachApiProductRequest) SetApiProductId(v string) *DetachApiProductRequest {
	s.ApiProductId = &v
	return s
}

func (s *DetachApiProductRequest) SetApis(v []*DetachApiProductRequestApis) *DetachApiProductRequest {
	s.Apis = v
	return s
}

func (s *DetachApiProductRequest) SetSecurityToken(v string) *DetachApiProductRequest {
	s.SecurityToken = &v
	return s
}

func (s *DetachApiProductRequest) Validate() error {
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

type DetachApiProductRequestApis struct {
	// The API ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ba84c55eca46488598da17c0609f3ead
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The environment to which the API is published. Valid values:
	//
	// 	- **RELEASE**: the production environment
	//
	// 	- **PRE**: the pre-release environment
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

func (s DetachApiProductRequestApis) String() string {
	return dara.Prettify(s)
}

func (s DetachApiProductRequestApis) GoString() string {
	return s.String()
}

func (s *DetachApiProductRequestApis) GetApiId() *string {
	return s.ApiId
}

func (s *DetachApiProductRequestApis) GetStageName() *string {
	return s.StageName
}

func (s *DetachApiProductRequestApis) SetApiId(v string) *DetachApiProductRequestApis {
	s.ApiId = &v
	return s
}

func (s *DetachApiProductRequestApis) SetStageName(v string) *DetachApiProductRequestApis {
	s.StageName = &v
	return s
}

func (s *DetachApiProductRequestApis) Validate() error {
	return dara.Validate(s)
}
