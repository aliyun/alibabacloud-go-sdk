// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAScriptsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAScriptIds(v []*string) *ListAScriptsRequest
	GetAScriptIds() []*string
	SetAScriptNames(v []*string) *ListAScriptsRequest
	GetAScriptNames() []*string
	SetListenerIds(v []*string) *ListAScriptsRequest
	GetListenerIds() []*string
	SetMaxResults(v int32) *ListAScriptsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAScriptsRequest
	GetNextToken() *string
}

type ListAScriptsRequest struct {
	// The AScript rule IDs. You can specify at most 20 IDs in each call.
	AScriptIds []*string `json:"AScriptIds,omitempty" xml:"AScriptIds,omitempty" type:"Repeated"`
	// The AScript rule names. You can specify at most 10 names in each call.
	AScriptNames []*string `json:"AScriptNames,omitempty" xml:"AScriptNames,omitempty" type:"Repeated"`
	// The listener IDs. You can specify at most 20 listener IDs in each call.
	ListenerIds []*string `json:"ListenerIds,omitempty" xml:"ListenerIds,omitempty" type:"Repeated"`
	// The maximum number of entries to return.
	//
	// Valid values: **1*	- to **100**.
	//
	// Default value: **20**. If you do not specify this parameter, the default value is used.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.****
	//
	// 	- You must specify the token that is obtained from the previous query as the value of **NextToken**.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListAScriptsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAScriptsRequest) GoString() string {
	return s.String()
}

func (s *ListAScriptsRequest) GetAScriptIds() []*string {
	return s.AScriptIds
}

func (s *ListAScriptsRequest) GetAScriptNames() []*string {
	return s.AScriptNames
}

func (s *ListAScriptsRequest) GetListenerIds() []*string {
	return s.ListenerIds
}

func (s *ListAScriptsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAScriptsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAScriptsRequest) SetAScriptIds(v []*string) *ListAScriptsRequest {
	s.AScriptIds = v
	return s
}

func (s *ListAScriptsRequest) SetAScriptNames(v []*string) *ListAScriptsRequest {
	s.AScriptNames = v
	return s
}

func (s *ListAScriptsRequest) SetListenerIds(v []*string) *ListAScriptsRequest {
	s.ListenerIds = v
	return s
}

func (s *ListAScriptsRequest) SetMaxResults(v int32) *ListAScriptsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAScriptsRequest) SetNextToken(v string) *ListAScriptsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAScriptsRequest) Validate() error {
	return dara.Validate(s)
}
