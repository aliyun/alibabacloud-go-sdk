// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHighCodeDeployResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *HighCodeDeployResponseBody
	GetData() *string
	SetErrorCode(v string) *HighCodeDeployResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *HighCodeDeployResponseBody
	GetErrorMsg() *string
	SetSuccess(v bool) *HighCodeDeployResponseBody
	GetSuccess() *bool
}

type HighCodeDeployResponseBody struct {
	// example:
	//
	// {\"key\": \"value\"}
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 404
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// 用户不存在
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Success  *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HighCodeDeployResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HighCodeDeployResponseBody) GoString() string {
	return s.String()
}

func (s *HighCodeDeployResponseBody) GetData() *string {
	return s.Data
}

func (s *HighCodeDeployResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *HighCodeDeployResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *HighCodeDeployResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HighCodeDeployResponseBody) SetData(v string) *HighCodeDeployResponseBody {
	s.Data = &v
	return s
}

func (s *HighCodeDeployResponseBody) SetErrorCode(v string) *HighCodeDeployResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *HighCodeDeployResponseBody) SetErrorMsg(v string) *HighCodeDeployResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *HighCodeDeployResponseBody) SetSuccess(v bool) *HighCodeDeployResponseBody {
	s.Success = &v
	return s
}

func (s *HighCodeDeployResponseBody) Validate() error {
	return dara.Validate(s)
}
