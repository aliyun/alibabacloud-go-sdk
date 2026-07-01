// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRcsSignMenuByVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRcsMenuVersion(v string) *QueryRcsSignMenuByVersionRequest
	GetRcsMenuVersion() *string
	SetSignName(v string) *QueryRcsSignMenuByVersionRequest
	GetSignName() *string
}

type QueryRcsSignMenuByVersionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	RcsMenuVersion *string `json:"RcsMenuVersion,omitempty" xml:"RcsMenuVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s QueryRcsSignMenuByVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRcsSignMenuByVersionRequest) GoString() string {
	return s.String()
}

func (s *QueryRcsSignMenuByVersionRequest) GetRcsMenuVersion() *string {
	return s.RcsMenuVersion
}

func (s *QueryRcsSignMenuByVersionRequest) GetSignName() *string {
	return s.SignName
}

func (s *QueryRcsSignMenuByVersionRequest) SetRcsMenuVersion(v string) *QueryRcsSignMenuByVersionRequest {
	s.RcsMenuVersion = &v
	return s
}

func (s *QueryRcsSignMenuByVersionRequest) SetSignName(v string) *QueryRcsSignMenuByVersionRequest {
	s.SignName = &v
	return s
}

func (s *QueryRcsSignMenuByVersionRequest) Validate() error {
	return dara.Validate(s)
}
