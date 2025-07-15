// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportOASRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthType(v string) *ImportOASRequest
	GetAuthType() *string
	SetBackendName(v string) *ImportOASRequest
	GetBackendName() *string
	SetData(v string) *ImportOASRequest
	GetData() *string
	SetGroupId(v string) *ImportOASRequest
	GetGroupId() *string
	SetIgnoreWarning(v bool) *ImportOASRequest
	GetIgnoreWarning() *bool
	SetOASVersion(v string) *ImportOASRequest
	GetOASVersion() *string
	SetOverwrite(v bool) *ImportOASRequest
	GetOverwrite() *bool
	SetRequestMode(v string) *ImportOASRequest
	GetRequestMode() *string
	SetSecurityToken(v string) *ImportOASRequest
	GetSecurityToken() *string
	SetSkipDryRun(v bool) *ImportOASRequest
	GetSkipDryRun() *bool
}

type ImportOASRequest struct {
	// The security authentication method of the API. Valid values:
	//
	// 	- **APP: Only authorized applications can call the API.**
	//
	// 	- **ANONYMOUS: The API can be anonymously called. In this mode, you must take note of the following rules:**
	//
	//     	- All users who have obtained the API service information can call this API. API Gateway does not authenticate callers and cannot set user-specific throttling policies. If you make this API public, set API-specific throttling policies.
	//
	// example:
	//
	// APP
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The name of the backend service.
	//
	// example:
	//
	// testBackendService
	BackendName *string `json:"BackendName,omitempty" xml:"BackendName,omitempty"`
	// The OAS-compliant text file or OSS object URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// swagger: "2.0"
	//
	// info:
	//
	//   version: "1.0.0"
	//
	//   title: "Swagger Petstore 2.0"
	//
	// basePath: "/"
	//
	// schemes:
	//
	// - "https"
	//
	// - "http"
	//
	// paths:
	//
	//   /pet/findByStatus:
	//
	//     get:
	//
	//       tags:
	//
	//       - "pet"
	//
	//       summary: "Finds Pets by status"
	//
	//       operationId: "findPetsByStatus"
	//
	//       parameters:
	//
	//       - name: "status"
	//
	//         in: "query"
	//
	//         required: true
	//
	//         type: "array"
	//
	//         items:
	//
	//           type: "string"
	//
	//           enum:
	//
	//           - "available"
	//
	//           - "pending"
	//
	//           - "sold"
	//
	//           default: "available"
	//
	//         collectionFormat: "multi"
	//
	//       responses:
	//
	//         "200":
	//
	//           description: "successful operation"
	//
	//           schema:
	//
	//             type: "array"
	//
	//             items:
	//
	//               $ref: "#/definitions/Pet"
	//
	//         "400":
	//
	//           description: "Invalid status value"
	//
	// definitions:
	//
	//   Category:
	//
	//     type: "object"
	//
	//     properties:
	//
	//       id:
	//
	//         type: "integer"
	//
	//         format: "int64"
	//
	//       name:
	//
	//         type: "string"
	//
	//   Tag:
	//
	//     type: "object"
	//
	//     properties:
	//
	//       id:
	//
	//         type: "integer"
	//
	//         format: "int64"
	//
	//       name:
	//
	//         type: "string"
	//
	//   Pet:
	//
	//     type: "object"
	//
	//     required:
	//
	//     - "name"
	//
	//     - "photoUrls"
	//
	//     properties:
	//
	//       id:
	//
	//         type: "integer"
	//
	//         format: "int64"
	//
	//       category:
	//
	//         $ref: "#/definitions/Category"
	//
	//       name:
	//
	//         type: "string"
	//
	//         example: "doggie"
	//
	//       photoUrls:
	//
	//         type: "array"
	//
	//         items:
	//
	//           type: "string"
	//
	//       tags:
	//
	//         type: "array"
	//
	//         items:
	//
	//           $ref: "#/definitions/Tag"
	//
	//       status:
	//
	//         type: "string"
	//
	//         description: "pet status in the store"
	//
	//         enum:
	//
	//         - "available"
	//
	//         - "pending"
	//
	//         - "sold"
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the API group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 08ae4aa0f95e4321849ee57f4e0b3077
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Specifies whether to ignore alerts.
	//
	// example:
	//
	// true
	IgnoreWarning *bool `json:"IgnoreWarning,omitempty" xml:"IgnoreWarning,omitempty"`
	// The OAS version.
	//
	// example:
	//
	// OAS2
	OASVersion *string `json:"OASVersion,omitempty" xml:"OASVersion,omitempty"`
	// Specifies whether to overwrite an existing API.
	//
	// If an existing API has the same HTTP request type and backend request path as the API to be imported, the existing API is overwritten.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Overwrite *bool `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// The request mode. Valid values:
	//
	// 	- MAPPING: Parameters are mapped. Unknown parameters are filtered out.
	//
	// 	- PASSTHROUGH: Parameters are passed through.
	//
	// example:
	//
	// PASSTHROUGH
	RequestMode   *string `json:"RequestMode,omitempty" xml:"RequestMode,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// Specifies whether to directly import the API without performing a precheck.
	//
	// example:
	//
	// true
	SkipDryRun *bool `json:"SkipDryRun,omitempty" xml:"SkipDryRun,omitempty"`
}

func (s ImportOASRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportOASRequest) GoString() string {
	return s.String()
}

func (s *ImportOASRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *ImportOASRequest) GetBackendName() *string {
	return s.BackendName
}

func (s *ImportOASRequest) GetData() *string {
	return s.Data
}

func (s *ImportOASRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ImportOASRequest) GetIgnoreWarning() *bool {
	return s.IgnoreWarning
}

func (s *ImportOASRequest) GetOASVersion() *string {
	return s.OASVersion
}

func (s *ImportOASRequest) GetOverwrite() *bool {
	return s.Overwrite
}

func (s *ImportOASRequest) GetRequestMode() *string {
	return s.RequestMode
}

func (s *ImportOASRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ImportOASRequest) GetSkipDryRun() *bool {
	return s.SkipDryRun
}

func (s *ImportOASRequest) SetAuthType(v string) *ImportOASRequest {
	s.AuthType = &v
	return s
}

func (s *ImportOASRequest) SetBackendName(v string) *ImportOASRequest {
	s.BackendName = &v
	return s
}

func (s *ImportOASRequest) SetData(v string) *ImportOASRequest {
	s.Data = &v
	return s
}

func (s *ImportOASRequest) SetGroupId(v string) *ImportOASRequest {
	s.GroupId = &v
	return s
}

func (s *ImportOASRequest) SetIgnoreWarning(v bool) *ImportOASRequest {
	s.IgnoreWarning = &v
	return s
}

func (s *ImportOASRequest) SetOASVersion(v string) *ImportOASRequest {
	s.OASVersion = &v
	return s
}

func (s *ImportOASRequest) SetOverwrite(v bool) *ImportOASRequest {
	s.Overwrite = &v
	return s
}

func (s *ImportOASRequest) SetRequestMode(v string) *ImportOASRequest {
	s.RequestMode = &v
	return s
}

func (s *ImportOASRequest) SetSecurityToken(v string) *ImportOASRequest {
	s.SecurityToken = &v
	return s
}

func (s *ImportOASRequest) SetSkipDryRun(v bool) *ImportOASRequest {
	s.SkipDryRun = &v
	return s
}

func (s *ImportOASRequest) Validate() error {
	return dara.Validate(s)
}
