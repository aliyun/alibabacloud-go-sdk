// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetUserRequest
	GetClusterId() *string
	SetUserName(v string) *GetUserRequest
	GetUserName() *string
}

type GetUserRequest struct {
	// The cluster ID.
	//
	// You can call [ListClusters](https://help.aliyun.com/document_detail/87116.html) to obtain the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The username.
	//
	// This parameter is required.
	//
	// example:
	//
	// testuser
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetUserRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *GetUserRequest) SetClusterId(v string) *GetUserRequest {
	s.ClusterId = &v
	return s
}

func (s *GetUserRequest) SetUserName(v string) *GetUserRequest {
	s.UserName = &v
	return s
}

func (s *GetUserRequest) Validate() error {
	return dara.Validate(s)
}
