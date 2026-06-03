// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSessionListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBeginDate(v string) *GetSessionListResponseBody
	GetBeginDate() *string
	SetEndDate(v string) *GetSessionListResponseBody
	GetEndDate() *string
	SetIncomplete(v string) *GetSessionListResponseBody
	GetIncomplete() *string
	SetPageNumber(v int32) *GetSessionListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *GetSessionListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetSessionListResponseBody
	GetRequestId() *string
	SetResults(v []*GetSessionListResponseBodyResults) *GetSessionListResponseBody
	GetResults() []*GetSessionListResponseBodyResults
	SetTotalCount(v int64) *GetSessionListResponseBody
	GetTotalCount() *int64
}

type GetSessionListResponseBody struct {
	// example:
	//
	// 2019-06-06 00:00:00
	BeginDate *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// example:
	//
	// 2019-06-06 23:59:59
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// true
	Incomplete *string `json:"Incomplete,omitempty" xml:"Incomplete,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1B217656-2267-4FBF-B26C-CDD201BDC3B8
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*GetSessionListResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// example:
	//
	// 10000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetSessionListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSessionListResponseBody) GoString() string {
	return s.String()
}

func (s *GetSessionListResponseBody) GetBeginDate() *string {
	return s.BeginDate
}

func (s *GetSessionListResponseBody) GetEndDate() *string {
	return s.EndDate
}

func (s *GetSessionListResponseBody) GetIncomplete() *string {
	return s.Incomplete
}

func (s *GetSessionListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetSessionListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSessionListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSessionListResponseBody) GetResults() []*GetSessionListResponseBodyResults {
	return s.Results
}

func (s *GetSessionListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetSessionListResponseBody) SetBeginDate(v string) *GetSessionListResponseBody {
	s.BeginDate = &v
	return s
}

func (s *GetSessionListResponseBody) SetEndDate(v string) *GetSessionListResponseBody {
	s.EndDate = &v
	return s
}

func (s *GetSessionListResponseBody) SetIncomplete(v string) *GetSessionListResponseBody {
	s.Incomplete = &v
	return s
}

func (s *GetSessionListResponseBody) SetPageNumber(v int32) *GetSessionListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetSessionListResponseBody) SetPageSize(v int32) *GetSessionListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetSessionListResponseBody) SetRequestId(v string) *GetSessionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSessionListResponseBody) SetResults(v []*GetSessionListResponseBodyResults) *GetSessionListResponseBody {
	s.Results = v
	return s
}

func (s *GetSessionListResponseBody) SetTotalCount(v int64) *GetSessionListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetSessionListResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSessionListResponseBodyResults struct {
	// example:
	//
	// 0
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// 2019-06-06 00:00:00
	CaptureTime *string `json:"CaptureTime,omitempty" xml:"CaptureTime,omitempty"`
	// example:
	//
	// 192.168.XX.XX
	ClientIp      *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ClientIpAlias *string `json:"ClientIpAlias,omitempty" xml:"ClientIpAlias,omitempty"`
	// example:
	//
	// 00163E06****
	ClientMac *string `json:"ClientMac,omitempty" xml:"ClientMac,omitempty"`
	// example:
	//
	// administrator
	ClientOsUser *string `json:"ClientOsUser,omitempty" xml:"ClientOsUser,omitempty"`
	// example:
	//
	// 15629
	ClientPort *int32 `json:"ClientPort,omitempty" xml:"ClientPort,omitempty"`
	// example:
	//
	// navicat
	ClientProgram *string `json:"ClientProgram,omitempty" xml:"ClientProgram,omitempty"`
	// example:
	//
	// 1
	DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// example:
	//
	// root
	DbUser *string `json:"DbUser,omitempty" xml:"DbUser,omitempty"`
	// example:
	//
	// 0
	LoginCode *int32 `json:"LoginCode,omitempty" xml:"LoginCode,omitempty"`
	// example:
	//
	// Access denied for user \\"root\\"@\\"192.168.XX.XX\\"(using password: YES)
	LoginMessage *string `json:"LoginMessage,omitempty" xml:"LoginMessage,omitempty"`
	// example:
	//
	// 192.168.XX.XX
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// example:
	//
	// 00163E06****
	ServerMac *string `json:"ServerMac,omitempty" xml:"ServerMac,omitempty"`
	// example:
	//
	// 3306
	ServerPort *int32 `json:"ServerPort,omitempty" xml:"ServerPort,omitempty"`
	// example:
	//
	// 3011610850021000000
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GetSessionListResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s GetSessionListResponseBodyResults) GoString() string {
	return s.String()
}

func (s *GetSessionListResponseBodyResults) GetAction() *string {
	return s.Action
}

func (s *GetSessionListResponseBodyResults) GetCaptureTime() *string {
	return s.CaptureTime
}

func (s *GetSessionListResponseBodyResults) GetClientIp() *string {
	return s.ClientIp
}

func (s *GetSessionListResponseBodyResults) GetClientIpAlias() *string {
	return s.ClientIpAlias
}

func (s *GetSessionListResponseBodyResults) GetClientMac() *string {
	return s.ClientMac
}

func (s *GetSessionListResponseBodyResults) GetClientOsUser() *string {
	return s.ClientOsUser
}

func (s *GetSessionListResponseBodyResults) GetClientPort() *int32 {
	return s.ClientPort
}

func (s *GetSessionListResponseBodyResults) GetClientProgram() *string {
	return s.ClientProgram
}

func (s *GetSessionListResponseBodyResults) GetDbId() *int32 {
	return s.DbId
}

func (s *GetSessionListResponseBodyResults) GetDbUser() *string {
	return s.DbUser
}

func (s *GetSessionListResponseBodyResults) GetLoginCode() *int32 {
	return s.LoginCode
}

func (s *GetSessionListResponseBodyResults) GetLoginMessage() *string {
	return s.LoginMessage
}

func (s *GetSessionListResponseBodyResults) GetServerIp() *string {
	return s.ServerIp
}

func (s *GetSessionListResponseBodyResults) GetServerMac() *string {
	return s.ServerMac
}

func (s *GetSessionListResponseBodyResults) GetServerPort() *int32 {
	return s.ServerPort
}

func (s *GetSessionListResponseBodyResults) GetSessionId() *string {
	return s.SessionId
}

func (s *GetSessionListResponseBodyResults) SetAction(v string) *GetSessionListResponseBodyResults {
	s.Action = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetCaptureTime(v string) *GetSessionListResponseBodyResults {
	s.CaptureTime = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetClientIp(v string) *GetSessionListResponseBodyResults {
	s.ClientIp = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetClientIpAlias(v string) *GetSessionListResponseBodyResults {
	s.ClientIpAlias = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetClientMac(v string) *GetSessionListResponseBodyResults {
	s.ClientMac = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetClientOsUser(v string) *GetSessionListResponseBodyResults {
	s.ClientOsUser = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetClientPort(v int32) *GetSessionListResponseBodyResults {
	s.ClientPort = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetClientProgram(v string) *GetSessionListResponseBodyResults {
	s.ClientProgram = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetDbId(v int32) *GetSessionListResponseBodyResults {
	s.DbId = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetDbUser(v string) *GetSessionListResponseBodyResults {
	s.DbUser = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetLoginCode(v int32) *GetSessionListResponseBodyResults {
	s.LoginCode = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetLoginMessage(v string) *GetSessionListResponseBodyResults {
	s.LoginMessage = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetServerIp(v string) *GetSessionListResponseBodyResults {
	s.ServerIp = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetServerMac(v string) *GetSessionListResponseBodyResults {
	s.ServerMac = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetServerPort(v int32) *GetSessionListResponseBodyResults {
	s.ServerPort = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetSessionId(v string) *GetSessionListResponseBodyResults {
	s.SessionId = &v
	return s
}

func (s *GetSessionListResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
