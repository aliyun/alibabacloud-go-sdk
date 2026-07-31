// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSemanticJobLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetSemanticJobLogResponseBodyData) *GetSemanticJobLogResponseBody
	GetData() []*GetSemanticJobLogResponseBodyData
	SetRequestId(v string) *GetSemanticJobLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSemanticJobLogResponseBody
	GetSuccess() *bool
}

type GetSemanticJobLogResponseBody struct {
	// The list of log segments returned by the executor. The current POP contract does not expose sqlIndex or offset externally. Log segments are returned based on the default behavior of the operation.
	Data []*GetSemanticJobLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID. Used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 676271D6-53B4-57BE-89FA-72F7AE1418DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSemanticJobLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSemanticJobLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetSemanticJobLogResponseBody) GetData() []*GetSemanticJobLogResponseBodyData {
	return s.Data
}

func (s *GetSemanticJobLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSemanticJobLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSemanticJobLogResponseBody) SetData(v []*GetSemanticJobLogResponseBodyData) *GetSemanticJobLogResponseBody {
	s.Data = v
	return s
}

func (s *GetSemanticJobLogResponseBody) SetRequestId(v string) *GetSemanticJobLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSemanticJobLogResponseBody) SetSuccess(v bool) *GetSemanticJobLogResponseBody {
	s.Success = &v
	return s
}

func (s *GetSemanticJobLogResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSemanticJobLogResponseBodyData struct {
	// The raw log text returned in this response.
	//
	// example:
	//
	// semantic job started
	LogContent *string `json:"LogContent,omitempty" xml:"LogContent,omitempty"`
	// Indicates whether the current log segment has been read to the end. A value of true indicates that no more content follows this segment.
	LogEnd *bool `json:"LogEnd,omitempty" xml:"LogEnd,omitempty"`
}

func (s GetSemanticJobLogResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSemanticJobLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSemanticJobLogResponseBodyData) GetLogContent() *string {
	return s.LogContent
}

func (s *GetSemanticJobLogResponseBodyData) GetLogEnd() *bool {
	return s.LogEnd
}

func (s *GetSemanticJobLogResponseBodyData) SetLogContent(v string) *GetSemanticJobLogResponseBodyData {
	s.LogContent = &v
	return s
}

func (s *GetSemanticJobLogResponseBodyData) SetLogEnd(v bool) *GetSemanticJobLogResponseBodyData {
	s.LogEnd = &v
	return s
}

func (s *GetSemanticJobLogResponseBodyData) Validate() error {
	return dara.Validate(s)
}
