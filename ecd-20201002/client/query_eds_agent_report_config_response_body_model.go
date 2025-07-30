// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEdsAgentReportConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QueryEdsAgentReportConfigResponseBodyData) *QueryEdsAgentReportConfigResponseBody
	GetData() *QueryEdsAgentReportConfigResponseBodyData
	SetRequestId(v string) *QueryEdsAgentReportConfigResponseBody
	GetRequestId() *string
}

type QueryEdsAgentReportConfigResponseBody struct {
	Data      *QueryEdsAgentReportConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryEdsAgentReportConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryEdsAgentReportConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEdsAgentReportConfigResponseBody) GetData() *QueryEdsAgentReportConfigResponseBodyData {
	return s.Data
}

func (s *QueryEdsAgentReportConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryEdsAgentReportConfigResponseBody) SetData(v *QueryEdsAgentReportConfigResponseBodyData) *QueryEdsAgentReportConfigResponseBody {
	s.Data = v
	return s
}

func (s *QueryEdsAgentReportConfigResponseBody) SetRequestId(v string) *QueryEdsAgentReportConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEdsAgentReportConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryEdsAgentReportConfigResponseBodyData struct {
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
}

func (s QueryEdsAgentReportConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryEdsAgentReportConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryEdsAgentReportConfigResponseBodyData) GetConfig() *string {
	return s.Config
}

func (s *QueryEdsAgentReportConfigResponseBodyData) SetConfig(v string) *QueryEdsAgentReportConfigResponseBodyData {
	s.Config = &v
	return s
}

func (s *QueryEdsAgentReportConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
