// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *UpdatePackageRequest
	GetBody() *string
}

type UpdatePackageRequest struct {
	// The request body parameters.
	//
	// example:
	//
	// {
	//
	//     "add": {
	//
	//         "allowedProjectList": [
	//
	//             {
	//
	//                 "label": "2",
	//
	//                 "project": "project_name"
	//
	//             }
	//
	//         ],
	//
	//         "resourceList": {
	//
	//             "table": [
	//
	//                 {
	//
	//                     "name": "table_name",
	//
	//                     "actions": [
	//
	//                         "Describe",
	//
	//                         "Select"
	//
	//                     ]
	//
	//                 },
	//
	//                 {
	//
	//                     "name": "table_name",
	//
	//                     "actions": [
	//
	//                         "Describe",
	//
	//                         "Select"
	//
	//                     ]
	//
	//                 }
	//
	//             ],
	//
	//             "resource": [
	//
	//                 {
	//
	//                     "name": "",
	//
	//                     "actions": []
	//
	//                 },
	//
	//                 {
	//
	//                     "name": "",
	//
	//                     "actions": []
	//
	//                 }
	//
	//             ],
	//
	//             "function": [
	//
	//                 {
	//
	//                     "name": "",
	//
	//                     "actions": []
	//
	//                 },
	//
	//                 {
	//
	//                     "name": "",
	//
	//                     "actions": []
	//
	//                 }
	//
	//             ]
	//
	//         }
	//
	//     },
	//
	//     "remove": {
	//
	//         "allowedProjectList": [
	//
	//             {
	//
	//                 "project": "project_name"
	//
	//             },
	//
	//             {
	//
	//                 "project": "project_2"
	//
	//             }
	//
	//         ],
	//
	//         "resourceList": {
	//
	//             "table": [
	//
	//                 {
	//
	//                     "name": "table_name"
	//
	//                 },
	//
	//                 {
	//
	//                     "name": "table_name"
	//
	//                 }
	//
	//             ],
	//
	//             "resource": [
	//
	//                 {
	//
	//                     "name": ""
	//
	//                 },
	//
	//                 {
	//
	//                     "name": ""
	//
	//                 }
	//
	//             ],
	//
	//             "function": [
	//
	//                 {
	//
	//                     "name": ""
	//
	//                 },
	//
	//                 {
	//
	//                     "name": ""
	//
	//                 }
	//
	//             ]
	//
	//         }
	//
	//     }
	//
	// }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePackageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePackageRequest) GoString() string {
	return s.String()
}

func (s *UpdatePackageRequest) GetBody() *string {
	return s.Body
}

func (s *UpdatePackageRequest) SetBody(v string) *UpdatePackageRequest {
	s.Body = &v
	return s
}

func (s *UpdatePackageRequest) Validate() error {
	return dara.Validate(s)
}
