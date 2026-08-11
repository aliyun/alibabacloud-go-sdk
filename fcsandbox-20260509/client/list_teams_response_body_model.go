// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTeamsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTeamsResponseBody
	GetCode() *string
	SetMessage(v string) *ListTeamsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListTeamsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTeamsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTeamsResponseBody
	GetRequestId() *string
	SetTeams(v []*E2BTeam) *ListTeamsResponseBody
	GetTeams() []*E2BTeam
	SetTotal(v int32) *ListTeamsResponseBody
	GetTotal() *int32
}

type ListTeamsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 20. Minimum value: 1. Maximum value: 50.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 7ADFF8D8-D4BA-5F79-AD49-DDABFEA59B6C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of teams.
	Teams []*E2BTeam `json:"teams,omitempty" xml:"teams,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 65
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListTeamsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTeamsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTeamsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTeamsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTeamsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTeamsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTeamsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTeamsResponseBody) GetTeams() []*E2BTeam {
	return s.Teams
}

func (s *ListTeamsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListTeamsResponseBody) SetCode(v string) *ListTeamsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTeamsResponseBody) SetMessage(v string) *ListTeamsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTeamsResponseBody) SetPageNumber(v int32) *ListTeamsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTeamsResponseBody) SetPageSize(v int32) *ListTeamsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTeamsResponseBody) SetRequestId(v string) *ListTeamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTeamsResponseBody) SetTeams(v []*E2BTeam) *ListTeamsResponseBody {
	s.Teams = v
	return s
}

func (s *ListTeamsResponseBody) SetTotal(v int32) *ListTeamsResponseBody {
	s.Total = &v
	return s
}

func (s *ListTeamsResponseBody) Validate() error {
	if s.Teams != nil {
		for _, item := range s.Teams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
