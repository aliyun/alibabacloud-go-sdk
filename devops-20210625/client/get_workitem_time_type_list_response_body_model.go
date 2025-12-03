// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkitemTimeTypeListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetWorkitemTimeTypeListResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetWorkitemTimeTypeListResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *GetWorkitemTimeTypeListResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetWorkitemTimeTypeListResponseBody
	GetSuccess() *string
	SetTimeType(v []*GetWorkitemTimeTypeListResponseBodyTimeType) *GetWorkitemTimeTypeListResponseBody
	GetTimeType() []*GetWorkitemTimeTypeListResponseBodyTimeType
}

type GetWorkitemTimeTypeListResponseBody struct {
	// example:
	//
	// Invalid.IdNotFound
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// HC93CE1A-8D7A-13A9-8306-7465DE2E5C0F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success  *string                                        `json:"success,omitempty" xml:"success,omitempty"`
	TimeType []*GetWorkitemTimeTypeListResponseBodyTimeType `json:"timeType,omitempty" xml:"timeType,omitempty" type:"Repeated"`
}

func (s GetWorkitemTimeTypeListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkitemTimeTypeListResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkitemTimeTypeListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetWorkitemTimeTypeListResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetWorkitemTimeTypeListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkitemTimeTypeListResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetWorkitemTimeTypeListResponseBody) GetTimeType() []*GetWorkitemTimeTypeListResponseBodyTimeType {
	return s.TimeType
}

func (s *GetWorkitemTimeTypeListResponseBody) SetErrorCode(v string) *GetWorkitemTimeTypeListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetWorkitemTimeTypeListResponseBody) SetErrorMsg(v string) *GetWorkitemTimeTypeListResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetWorkitemTimeTypeListResponseBody) SetRequestId(v string) *GetWorkitemTimeTypeListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkitemTimeTypeListResponseBody) SetSuccess(v string) *GetWorkitemTimeTypeListResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkitemTimeTypeListResponseBody) SetTimeType(v []*GetWorkitemTimeTypeListResponseBodyTimeType) *GetWorkitemTimeTypeListResponseBody {
	s.TimeType = v
	return s
}

func (s *GetWorkitemTimeTypeListResponseBody) Validate() error {
	if s.TimeType != nil {
		for _, item := range s.TimeType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWorkitemTimeTypeListResponseBodyTimeType struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// deploy
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// 67fb001005aac8d3d2a3372416
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	Position *int64 `json:"position,omitempty" xml:"position,omitempty"`
}

func (s GetWorkitemTimeTypeListResponseBodyTimeType) String() string {
	return dara.Prettify(s)
}

func (s GetWorkitemTimeTypeListResponseBodyTimeType) GoString() string {
	return s.String()
}

func (s *GetWorkitemTimeTypeListResponseBodyTimeType) GetDescription() *string {
	return s.Description
}

func (s *GetWorkitemTimeTypeListResponseBodyTimeType) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetWorkitemTimeTypeListResponseBodyTimeType) GetIdentifier() *string {
	return s.Identifier
}

func (s *GetWorkitemTimeTypeListResponseBodyTimeType) GetName() *string {
	return s.Name
}

func (s *GetWorkitemTimeTypeListResponseBodyTimeType) GetPosition() *int64 {
	return s.Position
}

func (s *GetWorkitemTimeTypeListResponseBodyTimeType) SetDescription(v string) *GetWorkitemTimeTypeListResponseBodyTimeType {
	s.Description = &v
	return s
}

func (s *GetWorkitemTimeTypeListResponseBodyTimeType) SetDisplayName(v string) *GetWorkitemTimeTypeListResponseBodyTimeType {
	s.DisplayName = &v
	return s
}

func (s *GetWorkitemTimeTypeListResponseBodyTimeType) SetIdentifier(v string) *GetWorkitemTimeTypeListResponseBodyTimeType {
	s.Identifier = &v
	return s
}

func (s *GetWorkitemTimeTypeListResponseBodyTimeType) SetName(v string) *GetWorkitemTimeTypeListResponseBodyTimeType {
	s.Name = &v
	return s
}

func (s *GetWorkitemTimeTypeListResponseBodyTimeType) SetPosition(v int64) *GetWorkitemTimeTypeListResponseBodyTimeType {
	s.Position = &v
	return s
}

func (s *GetWorkitemTimeTypeListResponseBodyTimeType) Validate() error {
	return dara.Validate(s)
}
