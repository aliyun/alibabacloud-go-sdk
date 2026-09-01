// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTreeDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPlaybooks(v string) *QueryTreeDataResponseBody
	GetPlaybooks() *string
	SetRequestId(v string) *QueryTreeDataResponseBody
	GetRequestId() *string
}

type QueryTreeDataResponseBody struct {
	// A list of playbooks in JSON string format. The string contains the following fields:
	//
	// - **active**: Indicates whether the playbook is active. A value of **true*	- means the playbook is active. A value of **false*	- means the playbook is inactive.
	//
	// - **displayName**: The name of the playbook.
	//
	// - **playbookUuid**: The UUID of the playbook.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "playbook": {
	//
	//             "active": false,
	//
	//             "displayName": "test_playbook",
	//
	//             "playbookUuid": "09a20455-3d3a-424c-a1df-xxxxxx"
	//
	//         }
	//
	//     }
	//
	// ]
	Playbooks *string `json:"Playbooks,omitempty" xml:"Playbooks,omitempty"`
	// The ID of the request. Alibaba Cloud generates this unique ID for each request. Use this ID to troubleshoot issues.
	//
	// example:
	//
	// EF2ECA2D-D8E6-5021-BF5C-19DD6D52C5B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryTreeDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTreeDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTreeDataResponseBody) GetPlaybooks() *string {
	return s.Playbooks
}

func (s *QueryTreeDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTreeDataResponseBody) SetPlaybooks(v string) *QueryTreeDataResponseBody {
	s.Playbooks = &v
	return s
}

func (s *QueryTreeDataResponseBody) SetRequestId(v string) *QueryTreeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTreeDataResponseBody) Validate() error {
	return dara.Validate(s)
}
