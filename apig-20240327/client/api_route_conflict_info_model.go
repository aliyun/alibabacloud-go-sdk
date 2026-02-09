// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApiRouteConflictInfo interface {
	dara.Model
	String() string
	GoString() string
	SetConflicts(v []*ApiRouteConflictInfoConflicts) *ApiRouteConflictInfo
	GetConflicts() []*ApiRouteConflictInfoConflicts
	SetDomainInfo(v *ApiRouteConflictInfoDomainInfo) *ApiRouteConflictInfo
	GetDomainInfo() *ApiRouteConflictInfoDomainInfo
}

type ApiRouteConflictInfo struct {
	// The conflicts.
	Conflicts []*ApiRouteConflictInfoConflicts `json:"conflicts,omitempty" xml:"conflicts,omitempty" type:"Repeated"`
	// The conflicting routes.
	DomainInfo *ApiRouteConflictInfoDomainInfo `json:"domainInfo,omitempty" xml:"domainInfo,omitempty" type:"Struct"`
}

func (s ApiRouteConflictInfo) String() string {
	return dara.Prettify(s)
}

func (s ApiRouteConflictInfo) GoString() string {
	return s.String()
}

func (s *ApiRouteConflictInfo) GetConflicts() []*ApiRouteConflictInfoConflicts {
	return s.Conflicts
}

func (s *ApiRouteConflictInfo) GetDomainInfo() *ApiRouteConflictInfoDomainInfo {
	return s.DomainInfo
}

func (s *ApiRouteConflictInfo) SetConflicts(v []*ApiRouteConflictInfoConflicts) *ApiRouteConflictInfo {
	s.Conflicts = v
	return s
}

func (s *ApiRouteConflictInfo) SetDomainInfo(v *ApiRouteConflictInfoDomainInfo) *ApiRouteConflictInfo {
	s.DomainInfo = v
	return s
}

func (s *ApiRouteConflictInfo) Validate() error {
	if s.Conflicts != nil {
		for _, item := range s.Conflicts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DomainInfo != nil {
		if err := s.DomainInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ApiRouteConflictInfoConflicts struct {
	// The details about the conflicts.
	Details []*ApiRouteConflictInfoConflictsDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
	// For APIs, conflicts exist in the specific environment. If the conflict object is a route, ignore.
	EnvironmentInfo *ApiRouteConflictInfoConflictsEnvironmentInfo `json:"environmentInfo,omitempty" xml:"environmentInfo,omitempty" type:"Struct"`
	// The conflicting resource ID.
	//
	// example:
	//
	// api-crdclqllhtggm***
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The conflicting resource name.
	//
	// example:
	//
	// item-stock
	ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	// The type of the conflicting resource.
	//
	// Valid values:
	//
	// 	- RestApi
	//
	// 	- HttpApiRoute
	//
	// example:
	//
	// HttpApiRoute
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The route information.
	RouteInfo *ApiRouteConflictInfoConflictsRouteInfo `json:"routeInfo,omitempty" xml:"routeInfo,omitempty" type:"Struct"`
}

func (s ApiRouteConflictInfoConflicts) String() string {
	return dara.Prettify(s)
}

func (s ApiRouteConflictInfoConflicts) GoString() string {
	return s.String()
}

func (s *ApiRouteConflictInfoConflicts) GetDetails() []*ApiRouteConflictInfoConflictsDetails {
	return s.Details
}

func (s *ApiRouteConflictInfoConflicts) GetEnvironmentInfo() *ApiRouteConflictInfoConflictsEnvironmentInfo {
	return s.EnvironmentInfo
}

func (s *ApiRouteConflictInfoConflicts) GetResourceId() *string {
	return s.ResourceId
}

func (s *ApiRouteConflictInfoConflicts) GetResourceName() *string {
	return s.ResourceName
}

func (s *ApiRouteConflictInfoConflicts) GetResourceType() *string {
	return s.ResourceType
}

func (s *ApiRouteConflictInfoConflicts) GetRouteInfo() *ApiRouteConflictInfoConflictsRouteInfo {
	return s.RouteInfo
}

func (s *ApiRouteConflictInfoConflicts) SetDetails(v []*ApiRouteConflictInfoConflictsDetails) *ApiRouteConflictInfoConflicts {
	s.Details = v
	return s
}

func (s *ApiRouteConflictInfoConflicts) SetEnvironmentInfo(v *ApiRouteConflictInfoConflictsEnvironmentInfo) *ApiRouteConflictInfoConflicts {
	s.EnvironmentInfo = v
	return s
}

func (s *ApiRouteConflictInfoConflicts) SetResourceId(v string) *ApiRouteConflictInfoConflicts {
	s.ResourceId = &v
	return s
}

func (s *ApiRouteConflictInfoConflicts) SetResourceName(v string) *ApiRouteConflictInfoConflicts {
	s.ResourceName = &v
	return s
}

func (s *ApiRouteConflictInfoConflicts) SetResourceType(v string) *ApiRouteConflictInfoConflicts {
	s.ResourceType = &v
	return s
}

func (s *ApiRouteConflictInfoConflicts) SetRouteInfo(v *ApiRouteConflictInfoConflictsRouteInfo) *ApiRouteConflictInfoConflicts {
	s.RouteInfo = v
	return s
}

func (s *ApiRouteConflictInfoConflicts) Validate() error {
	if s.Details != nil {
		for _, item := range s.Details {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EnvironmentInfo != nil {
		if err := s.EnvironmentInfo.Validate(); err != nil {
			return err
		}
	}
	if s.RouteInfo != nil {
		if err := s.RouteInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ApiRouteConflictInfoConflictsDetails struct {
	// The matching rule information of the conflicting target.
	ConflictingMatch *ApiRouteConflictInfoConflictsDetailsConflictingMatch `json:"conflictingMatch,omitempty" xml:"conflictingMatch,omitempty" type:"Struct"`
	// The matching rule information of the object being detected.
	DetectedMatch *ApiRouteConflictInfoConflictsDetailsDetectedMatch `json:"detectedMatch,omitempty" xml:"detectedMatch,omitempty" type:"Struct"`
	// The conflict level. Valid values: Critical, Warning, and Informational.
	//
	// example:
	//
	// Critical
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
}

func (s ApiRouteConflictInfoConflictsDetails) String() string {
	return dara.Prettify(s)
}

func (s ApiRouteConflictInfoConflictsDetails) GoString() string {
	return s.String()
}

func (s *ApiRouteConflictInfoConflictsDetails) GetConflictingMatch() *ApiRouteConflictInfoConflictsDetailsConflictingMatch {
	return s.ConflictingMatch
}

func (s *ApiRouteConflictInfoConflictsDetails) GetDetectedMatch() *ApiRouteConflictInfoConflictsDetailsDetectedMatch {
	return s.DetectedMatch
}

func (s *ApiRouteConflictInfoConflictsDetails) GetLevel() *string {
	return s.Level
}

func (s *ApiRouteConflictInfoConflictsDetails) SetConflictingMatch(v *ApiRouteConflictInfoConflictsDetailsConflictingMatch) *ApiRouteConflictInfoConflictsDetails {
	s.ConflictingMatch = v
	return s
}

func (s *ApiRouteConflictInfoConflictsDetails) SetDetectedMatch(v *ApiRouteConflictInfoConflictsDetailsDetectedMatch) *ApiRouteConflictInfoConflictsDetails {
	s.DetectedMatch = v
	return s
}

func (s *ApiRouteConflictInfoConflictsDetails) SetLevel(v string) *ApiRouteConflictInfoConflictsDetails {
	s.Level = &v
	return s
}

func (s *ApiRouteConflictInfoConflictsDetails) Validate() error {
	if s.ConflictingMatch != nil {
		if err := s.ConflictingMatch.Validate(); err != nil {
			return err
		}
	}
	if s.DetectedMatch != nil {
		if err := s.DetectedMatch.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ApiRouteConflictInfoConflictsDetailsConflictingMatch struct {
	// The matching rule.
	Match *HttpRouteMatch `json:"match,omitempty" xml:"match,omitempty"`
	// The corresponding operation information if the conflicting target is an API.
	OperationInfo *ApiRouteConflictInfoConflictsDetailsConflictingMatchOperationInfo `json:"operationInfo,omitempty" xml:"operationInfo,omitempty" type:"Struct"`
}

func (s ApiRouteConflictInfoConflictsDetailsConflictingMatch) String() string {
	return dara.Prettify(s)
}

func (s ApiRouteConflictInfoConflictsDetailsConflictingMatch) GoString() string {
	return s.String()
}

func (s *ApiRouteConflictInfoConflictsDetailsConflictingMatch) GetMatch() *HttpRouteMatch {
	return s.Match
}

func (s *ApiRouteConflictInfoConflictsDetailsConflictingMatch) GetOperationInfo() *ApiRouteConflictInfoConflictsDetailsConflictingMatchOperationInfo {
	return s.OperationInfo
}

func (s *ApiRouteConflictInfoConflictsDetailsConflictingMatch) SetMatch(v *HttpRouteMatch) *ApiRouteConflictInfoConflictsDetailsConflictingMatch {
	s.Match = v
	return s
}

func (s *ApiRouteConflictInfoConflictsDetailsConflictingMatch) SetOperationInfo(v *ApiRouteConflictInfoConflictsDetailsConflictingMatchOperationInfo) *ApiRouteConflictInfoConflictsDetailsConflictingMatch {
	s.OperationInfo = v
	return s
}

func (s *ApiRouteConflictInfoConflictsDetailsConflictingMatch) Validate() error {
	if s.Match != nil {
		if err := s.Match.Validate(); err != nil {
			return err
		}
	}
	if s.OperationInfo != nil {
		if err := s.OperationInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ApiRouteConflictInfoConflictsDetailsConflictingMatchOperationInfo struct {
	// The operation name.
	//
	// example:
	//
	// GetItem
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The operation ID.
	//
	// example:
	//
	// op-cq35jadlhtgrv07***
	OperationId *string `json:"operationId,omitempty" xml:"operationId,omitempty"`
}

func (s ApiRouteConflictInfoConflictsDetailsConflictingMatchOperationInfo) String() string {
	return dara.Prettify(s)
}

func (s ApiRouteConflictInfoConflictsDetailsConflictingMatchOperationInfo) GoString() string {
	return s.String()
}

func (s *ApiRouteConflictInfoConflictsDetailsConflictingMatchOperationInfo) GetName() *string {
	return s.Name
}

func (s *ApiRouteConflictInfoConflictsDetailsConflictingMatchOperationInfo) GetOperationId() *string {
	return s.OperationId
}

func (s *ApiRouteConflictInfoConflictsDetailsConflictingMatchOperationInfo) SetName(v string) *ApiRouteConflictInfoConflictsDetailsConflictingMatchOperationInfo {
	s.Name = &v
	return s
}

func (s *ApiRouteConflictInfoConflictsDetailsConflictingMatchOperationInfo) SetOperationId(v string) *ApiRouteConflictInfoConflictsDetailsConflictingMatchOperationInfo {
	s.OperationId = &v
	return s
}

func (s *ApiRouteConflictInfoConflictsDetailsConflictingMatchOperationInfo) Validate() error {
	return dara.Validate(s)
}

type ApiRouteConflictInfoConflictsDetailsDetectedMatch struct {
	// The matching rule information of the object being detected.
	Match *HttpRouteMatch `json:"match,omitempty" xml:"match,omitempty"`
	// If the object is an API, the conflicting operation information needs to be returned.
	OperationInfo *ApiRouteConflictInfoConflictsDetailsDetectedMatchOperationInfo `json:"operationInfo,omitempty" xml:"operationInfo,omitempty" type:"Struct"`
}

func (s ApiRouteConflictInfoConflictsDetailsDetectedMatch) String() string {
	return dara.Prettify(s)
}

func (s ApiRouteConflictInfoConflictsDetailsDetectedMatch) GoString() string {
	return s.String()
}

func (s *ApiRouteConflictInfoConflictsDetailsDetectedMatch) GetMatch() *HttpRouteMatch {
	return s.Match
}

func (s *ApiRouteConflictInfoConflictsDetailsDetectedMatch) GetOperationInfo() *ApiRouteConflictInfoConflictsDetailsDetectedMatchOperationInfo {
	return s.OperationInfo
}

func (s *ApiRouteConflictInfoConflictsDetailsDetectedMatch) SetMatch(v *HttpRouteMatch) *ApiRouteConflictInfoConflictsDetailsDetectedMatch {
	s.Match = v
	return s
}

func (s *ApiRouteConflictInfoConflictsDetailsDetectedMatch) SetOperationInfo(v *ApiRouteConflictInfoConflictsDetailsDetectedMatchOperationInfo) *ApiRouteConflictInfoConflictsDetailsDetectedMatch {
	s.OperationInfo = v
	return s
}

func (s *ApiRouteConflictInfoConflictsDetailsDetectedMatch) Validate() error {
	if s.Match != nil {
		if err := s.Match.Validate(); err != nil {
			return err
		}
	}
	if s.OperationInfo != nil {
		if err := s.OperationInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ApiRouteConflictInfoConflictsDetailsDetectedMatchOperationInfo struct {
	// The operation name.
	//
	// example:
	//
	// GetItemV2
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The operation ID.
	//
	// example:
	//
	// op-cqf17dllhtgng1m**
	OperationId *string `json:"operationId,omitempty" xml:"operationId,omitempty"`
}

func (s ApiRouteConflictInfoConflictsDetailsDetectedMatchOperationInfo) String() string {
	return dara.Prettify(s)
}

func (s ApiRouteConflictInfoConflictsDetailsDetectedMatchOperationInfo) GoString() string {
	return s.String()
}

func (s *ApiRouteConflictInfoConflictsDetailsDetectedMatchOperationInfo) GetName() *string {
	return s.Name
}

func (s *ApiRouteConflictInfoConflictsDetailsDetectedMatchOperationInfo) GetOperationId() *string {
	return s.OperationId
}

func (s *ApiRouteConflictInfoConflictsDetailsDetectedMatchOperationInfo) SetName(v string) *ApiRouteConflictInfoConflictsDetailsDetectedMatchOperationInfo {
	s.Name = &v
	return s
}

func (s *ApiRouteConflictInfoConflictsDetailsDetectedMatchOperationInfo) SetOperationId(v string) *ApiRouteConflictInfoConflictsDetailsDetectedMatchOperationInfo {
	s.OperationId = &v
	return s
}

func (s *ApiRouteConflictInfoConflictsDetailsDetectedMatchOperationInfo) Validate() error {
	return dara.Validate(s)
}

type ApiRouteConflictInfoConflictsEnvironmentInfo struct {
	// The environment ID.
	//
	// example:
	//
	// env-cquqsollhtgidd***
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The environment name.
	//
	// example:
	//
	// itemcenter-dev
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ApiRouteConflictInfoConflictsEnvironmentInfo) String() string {
	return dara.Prettify(s)
}

func (s ApiRouteConflictInfoConflictsEnvironmentInfo) GoString() string {
	return s.String()
}

func (s *ApiRouteConflictInfoConflictsEnvironmentInfo) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ApiRouteConflictInfoConflictsEnvironmentInfo) GetName() *string {
	return s.Name
}

func (s *ApiRouteConflictInfoConflictsEnvironmentInfo) SetEnvironmentId(v string) *ApiRouteConflictInfoConflictsEnvironmentInfo {
	s.EnvironmentId = &v
	return s
}

func (s *ApiRouteConflictInfoConflictsEnvironmentInfo) SetName(v string) *ApiRouteConflictInfoConflictsEnvironmentInfo {
	s.Name = &v
	return s
}

func (s *ApiRouteConflictInfoConflictsEnvironmentInfo) Validate() error {
	return dara.Validate(s)
}

type ApiRouteConflictInfoConflictsRouteInfo struct {
	// The route name.
	//
	// example:
	//
	// itemcenter-route
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The route ID.
	//
	// example:
	//
	// hr-cr82undlhtgrlej***
	RouteId *string `json:"routeId,omitempty" xml:"routeId,omitempty"`
}

func (s ApiRouteConflictInfoConflictsRouteInfo) String() string {
	return dara.Prettify(s)
}

func (s ApiRouteConflictInfoConflictsRouteInfo) GoString() string {
	return s.String()
}

func (s *ApiRouteConflictInfoConflictsRouteInfo) GetName() *string {
	return s.Name
}

func (s *ApiRouteConflictInfoConflictsRouteInfo) GetRouteId() *string {
	return s.RouteId
}

func (s *ApiRouteConflictInfoConflictsRouteInfo) SetName(v string) *ApiRouteConflictInfoConflictsRouteInfo {
	s.Name = &v
	return s
}

func (s *ApiRouteConflictInfoConflictsRouteInfo) SetRouteId(v string) *ApiRouteConflictInfoConflictsRouteInfo {
	s.RouteId = &v
	return s
}

func (s *ApiRouteConflictInfoConflictsRouteInfo) Validate() error {
	return dara.Validate(s)
}

type ApiRouteConflictInfoDomainInfo struct {
	// The domain name ID.
	//
	// example:
	//
	// d-cqookcllhtgvof7e***
	DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	// The domain name.
	//
	// example:
	//
	// httpbin
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ApiRouteConflictInfoDomainInfo) String() string {
	return dara.Prettify(s)
}

func (s ApiRouteConflictInfoDomainInfo) GoString() string {
	return s.String()
}

func (s *ApiRouteConflictInfoDomainInfo) GetDomainId() *string {
	return s.DomainId
}

func (s *ApiRouteConflictInfoDomainInfo) GetName() *string {
	return s.Name
}

func (s *ApiRouteConflictInfoDomainInfo) SetDomainId(v string) *ApiRouteConflictInfoDomainInfo {
	s.DomainId = &v
	return s
}

func (s *ApiRouteConflictInfoDomainInfo) SetName(v string) *ApiRouteConflictInfoDomainInfo {
	s.Name = &v
	return s
}

func (s *ApiRouteConflictInfoDomainInfo) Validate() error {
	return dara.Validate(s)
}
