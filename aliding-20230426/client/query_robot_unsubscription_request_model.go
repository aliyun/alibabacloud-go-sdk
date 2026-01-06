// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRobotUnsubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *QueryRobotUnsubscriptionRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QueryRobotUnsubscriptionRequest
	GetPageSize() *int32
	SetRobotCode(v string) *QueryRobotUnsubscriptionRequest
	GetRobotCode() *string
	SetSceneCode(v string) *QueryRobotUnsubscriptionRequest
	GetSceneCode() *string
}

type QueryRobotUnsubscriptionRequest struct {
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// dingxxxxxxxxxxxxxxxxxx
	RobotCode *string `json:"RobotCode,omitempty" xml:"RobotCode,omitempty"`
	// example:
	//
	// wb62cz0x
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
}

func (s QueryRobotUnsubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRobotUnsubscriptionRequest) GoString() string {
	return s.String()
}

func (s *QueryRobotUnsubscriptionRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryRobotUnsubscriptionRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryRobotUnsubscriptionRequest) GetRobotCode() *string {
	return s.RobotCode
}

func (s *QueryRobotUnsubscriptionRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *QueryRobotUnsubscriptionRequest) SetPageNo(v int32) *QueryRobotUnsubscriptionRequest {
	s.PageNo = &v
	return s
}

func (s *QueryRobotUnsubscriptionRequest) SetPageSize(v int32) *QueryRobotUnsubscriptionRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRobotUnsubscriptionRequest) SetRobotCode(v string) *QueryRobotUnsubscriptionRequest {
	s.RobotCode = &v
	return s
}

func (s *QueryRobotUnsubscriptionRequest) SetSceneCode(v string) *QueryRobotUnsubscriptionRequest {
	s.SceneCode = &v
	return s
}

func (s *QueryRobotUnsubscriptionRequest) Validate() error {
	return dara.Validate(s)
}
