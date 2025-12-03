// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkitemRelationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetWorkitemRelationsResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetWorkitemRelationsResponseBody
	GetErrorMsg() *string
	SetRelationList(v []*GetWorkitemRelationsResponseBodyRelationList) *GetWorkitemRelationsResponseBody
	GetRelationList() []*GetWorkitemRelationsResponseBodyRelationList
	SetRequestId(v string) *GetWorkitemRelationsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWorkitemRelationsResponseBody
	GetSuccess() *bool
}

type GetWorkitemRelationsResponseBody struct {
	// example:
	//
	// InvalidGroup.IdNotFound
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error
	ErrorMsg     *string                                         `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	RelationList []*GetWorkitemRelationsResponseBodyRelationList `json:"relationList,omitempty" xml:"relationList,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetWorkitemRelationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkitemRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkitemRelationsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetWorkitemRelationsResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetWorkitemRelationsResponseBody) GetRelationList() []*GetWorkitemRelationsResponseBodyRelationList {
	return s.RelationList
}

func (s *GetWorkitemRelationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkitemRelationsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWorkitemRelationsResponseBody) SetErrorCode(v string) *GetWorkitemRelationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetWorkitemRelationsResponseBody) SetErrorMsg(v string) *GetWorkitemRelationsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetWorkitemRelationsResponseBody) SetRelationList(v []*GetWorkitemRelationsResponseBodyRelationList) *GetWorkitemRelationsResponseBody {
	s.RelationList = v
	return s
}

func (s *GetWorkitemRelationsResponseBody) SetRequestId(v string) *GetWorkitemRelationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkitemRelationsResponseBody) SetSuccess(v bool) *GetWorkitemRelationsResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkitemRelationsResponseBody) Validate() error {
	if s.RelationList != nil {
		for _, item := range s.RelationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWorkitemRelationsResponseBodyRelationList struct {
	// example:
	//
	// aliyun_1384605
	AssignedTo *string `json:"assignedTo,omitempty" xml:"assignedTo,omitempty"`
	// example:
	//
	// Req
	CategoryIdentifier *string `json:"categoryIdentifier,omitempty" xml:"categoryIdentifier,omitempty"`
	// example:
	//
	// 1667205617061
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1667205617089
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// deafe5f33xxxxx6a259d8dafd
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// example:
	//
	// 2b856dxxxxxxb61d93676255ba
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	Subject         *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// example:
	//
	// 9uy29901re573f561d69jn40
	WorkitemTypeIdentifier *string `json:"workitemTypeIdentifier,omitempty" xml:"workitemTypeIdentifier,omitempty"`
}

func (s GetWorkitemRelationsResponseBodyRelationList) String() string {
	return dara.Prettify(s)
}

func (s GetWorkitemRelationsResponseBodyRelationList) GoString() string {
	return s.String()
}

func (s *GetWorkitemRelationsResponseBodyRelationList) GetAssignedTo() *string {
	return s.AssignedTo
}

func (s *GetWorkitemRelationsResponseBodyRelationList) GetCategoryIdentifier() *string {
	return s.CategoryIdentifier
}

func (s *GetWorkitemRelationsResponseBodyRelationList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetWorkitemRelationsResponseBodyRelationList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetWorkitemRelationsResponseBodyRelationList) GetIdentifier() *string {
	return s.Identifier
}

func (s *GetWorkitemRelationsResponseBodyRelationList) GetSpaceIdentifier() *string {
	return s.SpaceIdentifier
}

func (s *GetWorkitemRelationsResponseBodyRelationList) GetSubject() *string {
	return s.Subject
}

func (s *GetWorkitemRelationsResponseBodyRelationList) GetWorkitemTypeIdentifier() *string {
	return s.WorkitemTypeIdentifier
}

func (s *GetWorkitemRelationsResponseBodyRelationList) SetAssignedTo(v string) *GetWorkitemRelationsResponseBodyRelationList {
	s.AssignedTo = &v
	return s
}

func (s *GetWorkitemRelationsResponseBodyRelationList) SetCategoryIdentifier(v string) *GetWorkitemRelationsResponseBodyRelationList {
	s.CategoryIdentifier = &v
	return s
}

func (s *GetWorkitemRelationsResponseBodyRelationList) SetGmtCreate(v string) *GetWorkitemRelationsResponseBodyRelationList {
	s.GmtCreate = &v
	return s
}

func (s *GetWorkitemRelationsResponseBodyRelationList) SetGmtModified(v string) *GetWorkitemRelationsResponseBodyRelationList {
	s.GmtModified = &v
	return s
}

func (s *GetWorkitemRelationsResponseBodyRelationList) SetIdentifier(v string) *GetWorkitemRelationsResponseBodyRelationList {
	s.Identifier = &v
	return s
}

func (s *GetWorkitemRelationsResponseBodyRelationList) SetSpaceIdentifier(v string) *GetWorkitemRelationsResponseBodyRelationList {
	s.SpaceIdentifier = &v
	return s
}

func (s *GetWorkitemRelationsResponseBodyRelationList) SetSubject(v string) *GetWorkitemRelationsResponseBodyRelationList {
	s.Subject = &v
	return s
}

func (s *GetWorkitemRelationsResponseBodyRelationList) SetWorkitemTypeIdentifier(v string) *GetWorkitemRelationsResponseBodyRelationList {
	s.WorkitemTypeIdentifier = &v
	return s
}

func (s *GetWorkitemRelationsResponseBodyRelationList) Validate() error {
	return dara.Validate(s)
}
