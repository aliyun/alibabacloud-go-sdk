// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAScriptsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAScriptIds(v []*CreateAScriptsResponseBodyAScriptIds) *CreateAScriptsResponseBody
	GetAScriptIds() []*CreateAScriptsResponseBodyAScriptIds
	SetJobId(v string) *CreateAScriptsResponseBody
	GetJobId() *string
	SetRequestId(v string) *CreateAScriptsResponseBody
	GetRequestId() *string
}

type CreateAScriptsResponseBody struct {
	// The AScript rule IDs.
	AScriptIds []*CreateAScriptsResponseBodyAScriptIds `json:"AScriptIds,omitempty" xml:"AScriptIds,omitempty" type:"Repeated"`
	// The asynchronous task ID.
	//
	// example:
	//
	// 5c607642-535e-4e06-9d77-df53049b****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BF0FE763-9603-558F-A55B-0F4B9A3E3C02
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAScriptsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAScriptsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAScriptsResponseBody) GetAScriptIds() []*CreateAScriptsResponseBodyAScriptIds {
	return s.AScriptIds
}

func (s *CreateAScriptsResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CreateAScriptsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAScriptsResponseBody) SetAScriptIds(v []*CreateAScriptsResponseBodyAScriptIds) *CreateAScriptsResponseBody {
	s.AScriptIds = v
	return s
}

func (s *CreateAScriptsResponseBody) SetJobId(v string) *CreateAScriptsResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateAScriptsResponseBody) SetRequestId(v string) *CreateAScriptsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAScriptsResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateAScriptsResponseBodyAScriptIds struct {
	// The AScript rule ID.
	//
	// example:
	//
	// as-xvq5igaa7uv6vr****
	AScriptId *string `json:"AScriptId,omitempty" xml:"AScriptId,omitempty"`
}

func (s CreateAScriptsResponseBodyAScriptIds) String() string {
	return dara.Prettify(s)
}

func (s CreateAScriptsResponseBodyAScriptIds) GoString() string {
	return s.String()
}

func (s *CreateAScriptsResponseBodyAScriptIds) GetAScriptId() *string {
	return s.AScriptId
}

func (s *CreateAScriptsResponseBodyAScriptIds) SetAScriptId(v string) *CreateAScriptsResponseBodyAScriptIds {
	s.AScriptId = &v
	return s
}

func (s *CreateAScriptsResponseBodyAScriptIds) Validate() error {
	return dara.Validate(s)
}
