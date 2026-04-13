// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGmtCreateTime(v string) *GetSessionResponseBody
	GetGmtCreateTime() *string
	SetGmtModified(v string) *GetSessionResponseBody
	GetGmtModified() *string
	SetOwnerId(v string) *GetSessionResponseBody
	GetOwnerId() *string
	SetRequestId(v string) *GetSessionResponseBody
	GetRequestId() *string
	SetSessionId(v string) *GetSessionResponseBody
	GetSessionId() *string
	SetTitle(v string) *GetSessionResponseBody
	GetTitle() *string
	SetUserId(v string) *GetSessionResponseBody
	GetUserId() *string
}

type GetSessionResponseBody struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1557******4904
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// ECB1F5******C1BF5223
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// se-j6p******dram6
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Title     *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 27******94904
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSessionResponseBody) GoString() string {
	return s.String()
}

func (s *GetSessionResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetSessionResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetSessionResponseBody) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSessionResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *GetSessionResponseBody) GetTitle() *string {
	return s.Title
}

func (s *GetSessionResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetSessionResponseBody) SetGmtCreateTime(v string) *GetSessionResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetSessionResponseBody) SetGmtModified(v string) *GetSessionResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetSessionResponseBody) SetOwnerId(v string) *GetSessionResponseBody {
	s.OwnerId = &v
	return s
}

func (s *GetSessionResponseBody) SetRequestId(v string) *GetSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSessionResponseBody) SetSessionId(v string) *GetSessionResponseBody {
	s.SessionId = &v
	return s
}

func (s *GetSessionResponseBody) SetTitle(v string) *GetSessionResponseBody {
	s.Title = &v
	return s
}

func (s *GetSessionResponseBody) SetUserId(v string) *GetSessionResponseBody {
	s.UserId = &v
	return s
}

func (s *GetSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
