// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPdpLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListPdpLogsResponseBody
	GetRequestId() *string
	SetResult(v *PdpConfigPageResult) *ListPdpLogsResponseBody
	GetResult() *PdpConfigPageResult
}

type ListPdpLogsResponseBody struct {
	RequestId *string              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *PdpConfigPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListPdpLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPdpLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPdpLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPdpLogsResponseBody) GetResult() *PdpConfigPageResult {
	return s.Result
}

func (s *ListPdpLogsResponseBody) SetRequestId(v string) *ListPdpLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPdpLogsResponseBody) SetResult(v *PdpConfigPageResult) *ListPdpLogsResponseBody {
	s.Result = v
	return s
}

func (s *ListPdpLogsResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}
