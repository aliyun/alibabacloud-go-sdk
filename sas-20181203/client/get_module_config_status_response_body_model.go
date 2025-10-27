// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModuleConfigStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetModuleConfigStatusResponseBodyData) *GetModuleConfigStatusResponseBody
	GetData() *GetModuleConfigStatusResponseBodyData
	SetRequestId(v string) *GetModuleConfigStatusResponseBody
	GetRequestId() *string
}

type GetModuleConfigStatusResponseBody struct {
	// The data returned.
	Data *GetModuleConfigStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 843E4805-****-7EE12FA8DBFD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetModuleConfigStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetModuleConfigStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetModuleConfigStatusResponseBody) GetData() *GetModuleConfigStatusResponseBodyData {
	return s.Data
}

func (s *GetModuleConfigStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModuleConfigStatusResponseBody) SetData(v *GetModuleConfigStatusResponseBodyData) *GetModuleConfigStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetModuleConfigStatusResponseBody) SetRequestId(v string) *GetModuleConfigStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModuleConfigStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetModuleConfigStatusResponseBodyData struct {
	// The check results of the service modules.
	ModuleConfigResults []*GetModuleConfigStatusResponseBodyDataModuleConfigResults `json:"ModuleConfigResults,omitempty" xml:"ModuleConfigResults,omitempty" type:"Repeated"`
}

func (s GetModuleConfigStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetModuleConfigStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetModuleConfigStatusResponseBodyData) GetModuleConfigResults() []*GetModuleConfigStatusResponseBodyDataModuleConfigResults {
	return s.ModuleConfigResults
}

func (s *GetModuleConfigStatusResponseBodyData) SetModuleConfigResults(v []*GetModuleConfigStatusResponseBodyDataModuleConfigResults) *GetModuleConfigStatusResponseBodyData {
	s.ModuleConfigResults = v
	return s
}

func (s *GetModuleConfigStatusResponseBodyData) Validate() error {
	if s.ModuleConfigResults != nil {
		for _, item := range s.ModuleConfigResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetModuleConfigStatusResponseBodyDataModuleConfigResults struct {
	// The name of the check item. Valid values:
	//
	// 	- **Ransom**: The anti-ransomware policy is enabled.
	//
	// 	- **WebLock**: The web tamper proofing feature is enabled.
	//
	// 	- **Rasp**: Applications are added to the application protection feature.
	//
	// 	- **Image**: The container images that can be scanned are specified.
	//
	// 	- **Virus**: The periodic virus scan policy is enabled.
	//
	// example:
	//
	// Ransom
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// Indicates whether the service module passed the status check. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Pass *bool `json:"Pass,omitempty" xml:"Pass,omitempty"`
}

func (s GetModuleConfigStatusResponseBodyDataModuleConfigResults) String() string {
	return dara.Prettify(s)
}

func (s GetModuleConfigStatusResponseBodyDataModuleConfigResults) GoString() string {
	return s.String()
}

func (s *GetModuleConfigStatusResponseBodyDataModuleConfigResults) GetModuleName() *string {
	return s.ModuleName
}

func (s *GetModuleConfigStatusResponseBodyDataModuleConfigResults) GetPass() *bool {
	return s.Pass
}

func (s *GetModuleConfigStatusResponseBodyDataModuleConfigResults) SetModuleName(v string) *GetModuleConfigStatusResponseBodyDataModuleConfigResults {
	s.ModuleName = &v
	return s
}

func (s *GetModuleConfigStatusResponseBodyDataModuleConfigResults) SetPass(v bool) *GetModuleConfigStatusResponseBodyDataModuleConfigResults {
	s.Pass = &v
	return s
}

func (s *GetModuleConfigStatusResponseBodyDataModuleConfigResults) Validate() error {
	return dara.Validate(s)
}
