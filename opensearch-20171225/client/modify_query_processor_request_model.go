// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyQueryProcessorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v interface{}) *ModifyQueryProcessorRequest
	GetBody() interface{}
	SetDryRun(v bool) *ModifyQueryProcessorRequest
	GetDryRun() *bool
}

type ModifyQueryProcessorRequest struct {
	// The request parameters.
	//
	// example:
	//
	// {
	//
	//     "domain": "GENERAL",
	//
	//     "category": "",
	//
	//     "processors": [
	//
	//         {
	//
	//             "name": "synonym",
	//
	//             "useSystemDictionary": true
	//
	//         },
	//
	//         {
	//
	//             "name": "stop_word",
	//
	//             "useSystemDictionary": true
	//
	//         }
	//
	//     ]
	//
	// }
	Body interface{} `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether the request is a dry run.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s ModifyQueryProcessorRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyQueryProcessorRequest) GoString() string {
	return s.String()
}

func (s *ModifyQueryProcessorRequest) GetBody() interface{} {
	return s.Body
}

func (s *ModifyQueryProcessorRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyQueryProcessorRequest) SetBody(v interface{}) *ModifyQueryProcessorRequest {
	s.Body = v
	return s
}

func (s *ModifyQueryProcessorRequest) SetDryRun(v bool) *ModifyQueryProcessorRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyQueryProcessorRequest) Validate() error {
	return dara.Validate(s)
}
