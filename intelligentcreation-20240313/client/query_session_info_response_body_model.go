// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySessionInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQueryResourceInfoList(v []*QuerySessionInfoResponseBodyQueryResourceInfoList) *QuerySessionInfoResponseBody
	GetQueryResourceInfoList() []*QuerySessionInfoResponseBodyQueryResourceInfoList
	SetRequestId(v string) *QuerySessionInfoResponseBody
	GetRequestId() *string
	SetTotal(v int64) *QuerySessionInfoResponseBody
	GetTotal() *int64
}

type QuerySessionInfoResponseBody struct {
	QueryResourceInfoList []*QuerySessionInfoResponseBodyQueryResourceInfoList `json:"queryResourceInfoList,omitempty" xml:"queryResourceInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 4D902811-B75C-5D1B-8882-D515F8E2F977
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 26
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s QuerySessionInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySessionInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySessionInfoResponseBody) GetQueryResourceInfoList() []*QuerySessionInfoResponseBodyQueryResourceInfoList {
	return s.QueryResourceInfoList
}

func (s *QuerySessionInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySessionInfoResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *QuerySessionInfoResponseBody) SetQueryResourceInfoList(v []*QuerySessionInfoResponseBodyQueryResourceInfoList) *QuerySessionInfoResponseBody {
	s.QueryResourceInfoList = v
	return s
}

func (s *QuerySessionInfoResponseBody) SetRequestId(v string) *QuerySessionInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySessionInfoResponseBody) SetTotal(v int64) *QuerySessionInfoResponseBody {
	s.Total = &v
	return s
}

func (s *QuerySessionInfoResponseBody) Validate() error {
	if s.QueryResourceInfoList != nil {
		for _, item := range s.QueryResourceInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuerySessionInfoResponseBodyQueryResourceInfoList struct {
	// example:
	//
	// a169e9ec18404edc9972afd80866dc97
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// FREE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QuerySessionInfoResponseBodyQueryResourceInfoList) String() string {
	return dara.Prettify(s)
}

func (s QuerySessionInfoResponseBodyQueryResourceInfoList) GoString() string {
	return s.String()
}

func (s *QuerySessionInfoResponseBodyQueryResourceInfoList) GetSessionId() *string {
	return s.SessionId
}

func (s *QuerySessionInfoResponseBodyQueryResourceInfoList) GetStatus() *string {
	return s.Status
}

func (s *QuerySessionInfoResponseBodyQueryResourceInfoList) SetSessionId(v string) *QuerySessionInfoResponseBodyQueryResourceInfoList {
	s.SessionId = &v
	return s
}

func (s *QuerySessionInfoResponseBodyQueryResourceInfoList) SetStatus(v string) *QuerySessionInfoResponseBodyQueryResourceInfoList {
	s.Status = &v
	return s
}

func (s *QuerySessionInfoResponseBodyQueryResourceInfoList) Validate() error {
	return dara.Validate(s)
}
