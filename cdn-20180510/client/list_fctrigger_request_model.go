// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFCTriggerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventMetaName(v string) *ListFCTriggerRequest
	GetEventMetaName() *string
	SetEventMetaVersion(v string) *ListFCTriggerRequest
	GetEventMetaVersion() *string
}

type ListFCTriggerRequest struct {
	// The name of the event. You can specify only one name.
	//
	// This parameter is required.
	//
	// example:
	//
	// LogFileCreated
	EventMetaName *string `json:"EventMetaName,omitempty" xml:"EventMetaName,omitempty"`
	// The version number of the event. You can specify only one version number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0.0
	EventMetaVersion *string `json:"EventMetaVersion,omitempty" xml:"EventMetaVersion,omitempty"`
}

func (s ListFCTriggerRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFCTriggerRequest) GoString() string {
	return s.String()
}

func (s *ListFCTriggerRequest) GetEventMetaName() *string {
	return s.EventMetaName
}

func (s *ListFCTriggerRequest) GetEventMetaVersion() *string {
	return s.EventMetaVersion
}

func (s *ListFCTriggerRequest) SetEventMetaName(v string) *ListFCTriggerRequest {
	s.EventMetaName = &v
	return s
}

func (s *ListFCTriggerRequest) SetEventMetaVersion(v string) *ListFCTriggerRequest {
	s.EventMetaVersion = &v
	return s
}

func (s *ListFCTriggerRequest) Validate() error {
	return dara.Validate(s)
}
