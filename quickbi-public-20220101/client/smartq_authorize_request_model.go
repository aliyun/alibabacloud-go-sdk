// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmartqAuthorizeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCubeIds(v string) *SmartqAuthorizeRequest
	GetCubeIds() *string
	SetExpireDay(v string) *SmartqAuthorizeRequest
	GetExpireDay() *string
	SetLlmCubeThemes(v string) *SmartqAuthorizeRequest
	GetLlmCubeThemes() *string
	SetLlmCubes(v string) *SmartqAuthorizeRequest
	GetLlmCubes() *string
	SetOperationType(v int32) *SmartqAuthorizeRequest
	GetOperationType() *int32
	SetUserIds(v string) *SmartqAuthorizeRequest
	GetUserIds() *string
}

type SmartqAuthorizeRequest struct {
	// An array of dataset IDs. Separate multiple IDs with commas.
	//
	// 	Notice: This parameter is converted to the corresponding Q\\&A resource ID for authorization. If a \\`cubeId\\` does not correspond to an existing Q\\&A resource, an error is reported that the Q\\&A resource does not exist. Ensure that the \\`cubeId\\` is correct.
	//
	// example:
	//
	// wasdasd*******1235235sd,ASDAS*********ASDAW123
	CubeIds *string `json:"CubeIds,omitempty" xml:"CubeIds,omitempty"`
	// The expiration time. The default is seven days.
	//
	// Format: 2099-12-31
	//
	// example:
	//
	// 2099-12-31
	ExpireDay *string `json:"ExpireDay,omitempty" xml:"ExpireDay,omitempty"`
	// An array of analysis subject IDs. Separate multiple IDs with commas.
	//
	// example:
	//
	// wasdasd*******1235235sd,ASDAS*********ASDAW123
	LlmCubeThemes *string `json:"LlmCubeThemes,omitempty" xml:"LlmCubeThemes,omitempty"`
	// An array of Q\\&A resource IDs. Separate multiple IDs with commas.
	//
	// example:
	//
	// wasdasd*******1235235sd,ASDAS*********ASDAW123
	LlmCubes *string `json:"LlmCubes,omitempty" xml:"LlmCubes,omitempty"`
	// The operation type. Valid values:
	//
	// - 0: Grant authorization
	//
	// - 1: Delete authorization
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	OperationType *int32 `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// An array of user IDs. Separate multiple IDs with commas.
	//
	// 	Notice: The number of user IDs × (the number of Q\\&A resources + the number of analysis subjects) in a single request cannot exceed 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// wasdasd*******1235235sd,ASDAS*********ASDAW123
	UserIds *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s SmartqAuthorizeRequest) String() string {
	return dara.Prettify(s)
}

func (s SmartqAuthorizeRequest) GoString() string {
	return s.String()
}

func (s *SmartqAuthorizeRequest) GetCubeIds() *string {
	return s.CubeIds
}

func (s *SmartqAuthorizeRequest) GetExpireDay() *string {
	return s.ExpireDay
}

func (s *SmartqAuthorizeRequest) GetLlmCubeThemes() *string {
	return s.LlmCubeThemes
}

func (s *SmartqAuthorizeRequest) GetLlmCubes() *string {
	return s.LlmCubes
}

func (s *SmartqAuthorizeRequest) GetOperationType() *int32 {
	return s.OperationType
}

func (s *SmartqAuthorizeRequest) GetUserIds() *string {
	return s.UserIds
}

func (s *SmartqAuthorizeRequest) SetCubeIds(v string) *SmartqAuthorizeRequest {
	s.CubeIds = &v
	return s
}

func (s *SmartqAuthorizeRequest) SetExpireDay(v string) *SmartqAuthorizeRequest {
	s.ExpireDay = &v
	return s
}

func (s *SmartqAuthorizeRequest) SetLlmCubeThemes(v string) *SmartqAuthorizeRequest {
	s.LlmCubeThemes = &v
	return s
}

func (s *SmartqAuthorizeRequest) SetLlmCubes(v string) *SmartqAuthorizeRequest {
	s.LlmCubes = &v
	return s
}

func (s *SmartqAuthorizeRequest) SetOperationType(v int32) *SmartqAuthorizeRequest {
	s.OperationType = &v
	return s
}

func (s *SmartqAuthorizeRequest) SetUserIds(v string) *SmartqAuthorizeRequest {
	s.UserIds = &v
	return s
}

func (s *SmartqAuthorizeRequest) Validate() error {
	return dara.Validate(s)
}
