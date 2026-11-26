// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryScriptHotWordListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *QueryScriptHotWordListResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *QueryScriptHotWordListResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryScriptHotWordListResponseBody
	GetRequestId() *string
	SetScriptHotWordList(v []*AICoachScriptHotWordItem) *QueryScriptHotWordListResponseBody
	GetScriptHotWordList() []*AICoachScriptHotWordItem
	SetSuccess(v bool) *QueryScriptHotWordListResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *QueryScriptHotWordListResponseBody
	GetTotal() *int32
}

type QueryScriptHotWordListResponseBody struct {
	ErrorCode         *string                     `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage      *string                     `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId         *string                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ScriptHotWordList []*AICoachScriptHotWordItem `json:"scriptHotWordList,omitempty" xml:"scriptHotWordList,omitempty" type:"Repeated"`
	Success           *bool                       `json:"success,omitempty" xml:"success,omitempty"`
	Total             *int32                      `json:"total,omitempty" xml:"total,omitempty"`
}

func (s QueryScriptHotWordListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryScriptHotWordListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryScriptHotWordListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryScriptHotWordListResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryScriptHotWordListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryScriptHotWordListResponseBody) GetScriptHotWordList() []*AICoachScriptHotWordItem {
	return s.ScriptHotWordList
}

func (s *QueryScriptHotWordListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryScriptHotWordListResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *QueryScriptHotWordListResponseBody) SetErrorCode(v string) *QueryScriptHotWordListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryScriptHotWordListResponseBody) SetErrorMessage(v string) *QueryScriptHotWordListResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryScriptHotWordListResponseBody) SetRequestId(v string) *QueryScriptHotWordListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryScriptHotWordListResponseBody) SetScriptHotWordList(v []*AICoachScriptHotWordItem) *QueryScriptHotWordListResponseBody {
	s.ScriptHotWordList = v
	return s
}

func (s *QueryScriptHotWordListResponseBody) SetSuccess(v bool) *QueryScriptHotWordListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryScriptHotWordListResponseBody) SetTotal(v int32) *QueryScriptHotWordListResponseBody {
	s.Total = &v
	return s
}

func (s *QueryScriptHotWordListResponseBody) Validate() error {
	if s.ScriptHotWordList != nil {
		for _, item := range s.ScriptHotWordList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
