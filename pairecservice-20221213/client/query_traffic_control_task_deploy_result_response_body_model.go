// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTrafficControlTaskDeployResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeployMessage(v string) *QueryTrafficControlTaskDeployResultResponseBody
	GetDeployMessage() *string
	SetDeployStatus(v string) *QueryTrafficControlTaskDeployResultResponseBody
	GetDeployStatus() *string
	SetDraftMessage(v string) *QueryTrafficControlTaskDeployResultResponseBody
	GetDraftMessage() *string
	SetDraftStatus(v string) *QueryTrafficControlTaskDeployResultResponseBody
	GetDraftStatus() *string
	SetPrepareMessage(v string) *QueryTrafficControlTaskDeployResultResponseBody
	GetPrepareMessage() *string
	SetPrepareStatus(v string) *QueryTrafficControlTaskDeployResultResponseBody
	GetPrepareStatus() *string
	SetRequestId(v string) *QueryTrafficControlTaskDeployResultResponseBody
	GetRequestId() *string
	SetStartMessage(v string) *QueryTrafficControlTaskDeployResultResponseBody
	GetStartMessage() *string
	SetStartStatus(v string) *QueryTrafficControlTaskDeployResultResponseBody
	GetStartStatus() *string
	SetTrafficControlTaskId(v string) *QueryTrafficControlTaskDeployResultResponseBody
	GetTrafficControlTaskId() *string
}

type QueryTrafficControlTaskDeployResultResponseBody struct {
	// example:
	//
	// deploy job draft success
	DeployMessage *string `json:"DeployMessage,omitempty" xml:"DeployMessage,omitempty"`
	// example:
	//
	// Success
	DeployStatus *string `json:"DeployStatus,omitempty" xml:"DeployStatus,omitempty"`
	// example:
	//
	// push draft success
	DraftMessage *string `json:"DraftMessage,omitempty" xml:"DraftMessage,omitempty"`
	// example:
	//
	// Success
	DraftStatus *string `json:"DraftStatus,omitempty" xml:"DraftStatus,omitempty"`
	// example:
	//
	// success
	PrepareMessage *string `json:"PrepareMessage,omitempty" xml:"PrepareMessage,omitempty"`
	// example:
	//
	// Success
	PrepareStatus *string `json:"PrepareStatus,omitempty" xml:"PrepareStatus,omitempty"`
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// start job success
	StartMessage *string `json:"StartMessage,omitempty" xml:"StartMessage,omitempty"`
	// example:
	//
	// Success
	StartStatus *string `json:"StartStatus,omitempty" xml:"StartStatus,omitempty"`
	// example:
	//
	// 3
	TrafficControlTaskId *string `json:"TrafficControlTaskId,omitempty" xml:"TrafficControlTaskId,omitempty"`
}

func (s QueryTrafficControlTaskDeployResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTrafficControlTaskDeployResultResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTrafficControlTaskDeployResultResponseBody) GetDeployMessage() *string {
	return s.DeployMessage
}

func (s *QueryTrafficControlTaskDeployResultResponseBody) GetDeployStatus() *string {
	return s.DeployStatus
}

func (s *QueryTrafficControlTaskDeployResultResponseBody) GetDraftMessage() *string {
	return s.DraftMessage
}

func (s *QueryTrafficControlTaskDeployResultResponseBody) GetDraftStatus() *string {
	return s.DraftStatus
}

func (s *QueryTrafficControlTaskDeployResultResponseBody) GetPrepareMessage() *string {
	return s.PrepareMessage
}

func (s *QueryTrafficControlTaskDeployResultResponseBody) GetPrepareStatus() *string {
	return s.PrepareStatus
}

func (s *QueryTrafficControlTaskDeployResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTrafficControlTaskDeployResultResponseBody) GetStartMessage() *string {
	return s.StartMessage
}

func (s *QueryTrafficControlTaskDeployResultResponseBody) GetStartStatus() *string {
	return s.StartStatus
}

func (s *QueryTrafficControlTaskDeployResultResponseBody) GetTrafficControlTaskId() *string {
	return s.TrafficControlTaskId
}

func (s *QueryTrafficControlTaskDeployResultResponseBody) SetDeployMessage(v string) *QueryTrafficControlTaskDeployResultResponseBody {
	s.DeployMessage = &v
	return s
}

func (s *QueryTrafficControlTaskDeployResultResponseBody) SetDeployStatus(v string) *QueryTrafficControlTaskDeployResultResponseBody {
	s.DeployStatus = &v
	return s
}

func (s *QueryTrafficControlTaskDeployResultResponseBody) SetDraftMessage(v string) *QueryTrafficControlTaskDeployResultResponseBody {
	s.DraftMessage = &v
	return s
}

func (s *QueryTrafficControlTaskDeployResultResponseBody) SetDraftStatus(v string) *QueryTrafficControlTaskDeployResultResponseBody {
	s.DraftStatus = &v
	return s
}

func (s *QueryTrafficControlTaskDeployResultResponseBody) SetPrepareMessage(v string) *QueryTrafficControlTaskDeployResultResponseBody {
	s.PrepareMessage = &v
	return s
}

func (s *QueryTrafficControlTaskDeployResultResponseBody) SetPrepareStatus(v string) *QueryTrafficControlTaskDeployResultResponseBody {
	s.PrepareStatus = &v
	return s
}

func (s *QueryTrafficControlTaskDeployResultResponseBody) SetRequestId(v string) *QueryTrafficControlTaskDeployResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTrafficControlTaskDeployResultResponseBody) SetStartMessage(v string) *QueryTrafficControlTaskDeployResultResponseBody {
	s.StartMessage = &v
	return s
}

func (s *QueryTrafficControlTaskDeployResultResponseBody) SetStartStatus(v string) *QueryTrafficControlTaskDeployResultResponseBody {
	s.StartStatus = &v
	return s
}

func (s *QueryTrafficControlTaskDeployResultResponseBody) SetTrafficControlTaskId(v string) *QueryTrafficControlTaskDeployResultResponseBody {
	s.TrafficControlTaskId = &v
	return s
}

func (s *QueryTrafficControlTaskDeployResultResponseBody) Validate() error {
	return dara.Validate(s)
}
