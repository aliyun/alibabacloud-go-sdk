// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectModelTierResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *UpdateProjectModelTierResponseBody
	GetData() *string
	SetErrorCode(v string) *UpdateProjectModelTierResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *UpdateProjectModelTierResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *UpdateProjectModelTierResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *UpdateProjectModelTierResponseBody
	GetRequestId() *string
}

type UpdateProjectModelTierResponseBody struct {
	// example:
	//
	// "data":{
	//
	// 		"data":"success",
	//
	// 		"requestId":"****"
	//
	// 	}
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// OBJECT_NOT_EXIST
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// this project is not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// 73207140-0FD5-588A-B11A-3CE093924196
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateProjectModelTierResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectModelTierResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectModelTierResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateProjectModelTierResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateProjectModelTierResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *UpdateProjectModelTierResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *UpdateProjectModelTierResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProjectModelTierResponseBody) SetData(v string) *UpdateProjectModelTierResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateProjectModelTierResponseBody) SetErrorCode(v string) *UpdateProjectModelTierResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateProjectModelTierResponseBody) SetErrorMsg(v string) *UpdateProjectModelTierResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateProjectModelTierResponseBody) SetHttpCode(v int32) *UpdateProjectModelTierResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateProjectModelTierResponseBody) SetRequestId(v string) *UpdateProjectModelTierResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectModelTierResponseBody) Validate() error {
	return dara.Validate(s)
}
