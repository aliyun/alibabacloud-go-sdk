// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpcutil "github.com/alibabacloud-go/tea-rpc-utils/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetAddressGeocodeRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty" require:"true"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty" require:"true"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty" require:"true"`
}

func (s GetAddressGeocodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAddressGeocodeRequest) GoString() string {
	return s.String()
}

func (s *GetAddressGeocodeRequest) SetServiceCode(v string) *GetAddressGeocodeRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetAddressGeocodeRequest) SetText(v string) *GetAddressGeocodeRequest {
	s.Text = &v
	return s
}

func (s *GetAddressGeocodeRequest) SetDefaultProvince(v string) *GetAddressGeocodeRequest {
	s.DefaultProvince = &v
	return s
}

func (s *GetAddressGeocodeRequest) SetDefaultCity(v string) *GetAddressGeocodeRequest {
	s.DefaultCity = &v
	return s
}

func (s *GetAddressGeocodeRequest) SetDefaultDistrict(v string) *GetAddressGeocodeRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *GetAddressGeocodeRequest) SetAppKey(v string) *GetAddressGeocodeRequest {
	s.AppKey = &v
	return s
}

type GetAddressGeocodeResponse struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s GetAddressGeocodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAddressGeocodeResponse) GoString() string {
	return s.String()
}

func (s *GetAddressGeocodeResponse) SetData(v string) *GetAddressGeocodeResponse {
	s.Data = &v
	return s
}

func (s *GetAddressGeocodeResponse) SetRequestId(v string) *GetAddressGeocodeResponse {
	s.RequestId = &v
	return s
}

type CompleteAddressRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty" require:"true"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty" require:"true"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty" require:"true"`
}

func (s CompleteAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s CompleteAddressRequest) GoString() string {
	return s.String()
}

func (s *CompleteAddressRequest) SetServiceCode(v string) *CompleteAddressRequest {
	s.ServiceCode = &v
	return s
}

func (s *CompleteAddressRequest) SetText(v string) *CompleteAddressRequest {
	s.Text = &v
	return s
}

func (s *CompleteAddressRequest) SetDefaultProvince(v string) *CompleteAddressRequest {
	s.DefaultProvince = &v
	return s
}

func (s *CompleteAddressRequest) SetDefaultCity(v string) *CompleteAddressRequest {
	s.DefaultCity = &v
	return s
}

func (s *CompleteAddressRequest) SetDefaultDistrict(v string) *CompleteAddressRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *CompleteAddressRequest) SetAppKey(v string) *CompleteAddressRequest {
	s.AppKey = &v
	return s
}

type CompleteAddressResponse struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s CompleteAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s CompleteAddressResponse) GoString() string {
	return s.String()
}

func (s *CompleteAddressResponse) SetData(v string) *CompleteAddressResponse {
	s.Data = &v
	return s
}

func (s *CompleteAddressResponse) SetRequestId(v string) *CompleteAddressResponse {
	s.RequestId = &v
	return s
}

type GetZipcodeRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty" require:"true"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty" require:"true"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty" require:"true"`
}

func (s GetZipcodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetZipcodeRequest) GoString() string {
	return s.String()
}

func (s *GetZipcodeRequest) SetServiceCode(v string) *GetZipcodeRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetZipcodeRequest) SetText(v string) *GetZipcodeRequest {
	s.Text = &v
	return s
}

func (s *GetZipcodeRequest) SetDefaultProvince(v string) *GetZipcodeRequest {
	s.DefaultProvince = &v
	return s
}

func (s *GetZipcodeRequest) SetDefaultCity(v string) *GetZipcodeRequest {
	s.DefaultCity = &v
	return s
}

func (s *GetZipcodeRequest) SetDefaultDistrict(v string) *GetZipcodeRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *GetZipcodeRequest) SetAppKey(v string) *GetZipcodeRequest {
	s.AppKey = &v
	return s
}

type GetZipcodeResponse struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s GetZipcodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetZipcodeResponse) GoString() string {
	return s.String()
}

func (s *GetZipcodeResponse) SetData(v string) *GetZipcodeResponse {
	s.Data = &v
	return s
}

func (s *GetZipcodeResponse) SetRequestId(v string) *GetZipcodeResponse {
	s.RequestId = &v
	return s
}

type ExtractPhoneRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty" require:"true"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty" require:"true"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty" require:"true"`
}

func (s ExtractPhoneRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtractPhoneRequest) GoString() string {
	return s.String()
}

func (s *ExtractPhoneRequest) SetServiceCode(v string) *ExtractPhoneRequest {
	s.ServiceCode = &v
	return s
}

func (s *ExtractPhoneRequest) SetText(v string) *ExtractPhoneRequest {
	s.Text = &v
	return s
}

func (s *ExtractPhoneRequest) SetDefaultProvince(v string) *ExtractPhoneRequest {
	s.DefaultProvince = &v
	return s
}

func (s *ExtractPhoneRequest) SetDefaultCity(v string) *ExtractPhoneRequest {
	s.DefaultCity = &v
	return s
}

func (s *ExtractPhoneRequest) SetDefaultDistrict(v string) *ExtractPhoneRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *ExtractPhoneRequest) SetAppKey(v string) *ExtractPhoneRequest {
	s.AppKey = &v
	return s
}

type ExtractPhoneResponse struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ExtractPhoneResponse) String() string {
	return tea.Prettify(s)
}

func (s ExtractPhoneResponse) GoString() string {
	return s.String()
}

func (s *ExtractPhoneResponse) SetData(v string) *ExtractPhoneResponse {
	s.Data = &v
	return s
}

func (s *ExtractPhoneResponse) SetRequestId(v string) *ExtractPhoneResponse {
	s.RequestId = &v
	return s
}

type ExtractNameRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty" require:"true"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty" require:"true"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty" require:"true"`
}

func (s ExtractNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtractNameRequest) GoString() string {
	return s.String()
}

func (s *ExtractNameRequest) SetServiceCode(v string) *ExtractNameRequest {
	s.ServiceCode = &v
	return s
}

func (s *ExtractNameRequest) SetText(v string) *ExtractNameRequest {
	s.Text = &v
	return s
}

func (s *ExtractNameRequest) SetDefaultProvince(v string) *ExtractNameRequest {
	s.DefaultProvince = &v
	return s
}

func (s *ExtractNameRequest) SetDefaultCity(v string) *ExtractNameRequest {
	s.DefaultCity = &v
	return s
}

func (s *ExtractNameRequest) SetDefaultDistrict(v string) *ExtractNameRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *ExtractNameRequest) SetAppKey(v string) *ExtractNameRequest {
	s.AppKey = &v
	return s
}

type ExtractNameResponse struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ExtractNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ExtractNameResponse) GoString() string {
	return s.String()
}

func (s *ExtractNameResponse) SetData(v string) *ExtractNameResponse {
	s.Data = &v
	return s
}

func (s *ExtractNameResponse) SetRequestId(v string) *ExtractNameResponse {
	s.RequestId = &v
	return s
}

type GetAddressDivisionCodeRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty" require:"true"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty" require:"true"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty" require:"true"`
}

func (s GetAddressDivisionCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAddressDivisionCodeRequest) GoString() string {
	return s.String()
}

func (s *GetAddressDivisionCodeRequest) SetServiceCode(v string) *GetAddressDivisionCodeRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetAddressDivisionCodeRequest) SetText(v string) *GetAddressDivisionCodeRequest {
	s.Text = &v
	return s
}

func (s *GetAddressDivisionCodeRequest) SetDefaultProvince(v string) *GetAddressDivisionCodeRequest {
	s.DefaultProvince = &v
	return s
}

func (s *GetAddressDivisionCodeRequest) SetDefaultCity(v string) *GetAddressDivisionCodeRequest {
	s.DefaultCity = &v
	return s
}

func (s *GetAddressDivisionCodeRequest) SetDefaultDistrict(v string) *GetAddressDivisionCodeRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *GetAddressDivisionCodeRequest) SetAppKey(v string) *GetAddressDivisionCodeRequest {
	s.AppKey = &v
	return s
}

type GetAddressDivisionCodeResponse struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s GetAddressDivisionCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAddressDivisionCodeResponse) GoString() string {
	return s.String()
}

func (s *GetAddressDivisionCodeResponse) SetData(v string) *GetAddressDivisionCodeResponse {
	s.Data = &v
	return s
}

func (s *GetAddressDivisionCodeResponse) SetRequestId(v string) *GetAddressDivisionCodeResponse {
	s.RequestId = &v
	return s
}

type ClassifyPOIRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty" require:"true"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty" require:"true"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty" require:"true"`
}

func (s ClassifyPOIRequest) String() string {
	return tea.Prettify(s)
}

func (s ClassifyPOIRequest) GoString() string {
	return s.String()
}

func (s *ClassifyPOIRequest) SetServiceCode(v string) *ClassifyPOIRequest {
	s.ServiceCode = &v
	return s
}

func (s *ClassifyPOIRequest) SetText(v string) *ClassifyPOIRequest {
	s.Text = &v
	return s
}

func (s *ClassifyPOIRequest) SetDefaultProvince(v string) *ClassifyPOIRequest {
	s.DefaultProvince = &v
	return s
}

func (s *ClassifyPOIRequest) SetDefaultCity(v string) *ClassifyPOIRequest {
	s.DefaultCity = &v
	return s
}

func (s *ClassifyPOIRequest) SetDefaultDistrict(v string) *ClassifyPOIRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *ClassifyPOIRequest) SetAppKey(v string) *ClassifyPOIRequest {
	s.AppKey = &v
	return s
}

type ClassifyPOIResponse struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ClassifyPOIResponse) String() string {
	return tea.Prettify(s)
}

func (s ClassifyPOIResponse) GoString() string {
	return s.String()
}

func (s *ClassifyPOIResponse) SetData(v string) *ClassifyPOIResponse {
	s.Data = &v
	return s
}

func (s *ClassifyPOIResponse) SetRequestId(v string) *ClassifyPOIResponse {
	s.RequestId = &v
	return s
}

type StructureAddressRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty" require:"true"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty" require:"true"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty" require:"true"`
}

func (s StructureAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s StructureAddressRequest) GoString() string {
	return s.String()
}

func (s *StructureAddressRequest) SetServiceCode(v string) *StructureAddressRequest {
	s.ServiceCode = &v
	return s
}

func (s *StructureAddressRequest) SetText(v string) *StructureAddressRequest {
	s.Text = &v
	return s
}

func (s *StructureAddressRequest) SetDefaultProvince(v string) *StructureAddressRequest {
	s.DefaultProvince = &v
	return s
}

func (s *StructureAddressRequest) SetDefaultCity(v string) *StructureAddressRequest {
	s.DefaultCity = &v
	return s
}

func (s *StructureAddressRequest) SetDefaultDistrict(v string) *StructureAddressRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *StructureAddressRequest) SetAppKey(v string) *StructureAddressRequest {
	s.AppKey = &v
	return s
}

type StructureAddressResponse struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s StructureAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s StructureAddressResponse) GoString() string {
	return s.String()
}

func (s *StructureAddressResponse) SetData(v string) *StructureAddressResponse {
	s.Data = &v
	return s
}

func (s *StructureAddressResponse) SetRequestId(v string) *StructureAddressResponse {
	s.RequestId = &v
	return s
}

type ExtractAddressRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty" require:"true"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty" require:"true"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty" require:"true"`
}

func (s ExtractAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtractAddressRequest) GoString() string {
	return s.String()
}

func (s *ExtractAddressRequest) SetServiceCode(v string) *ExtractAddressRequest {
	s.ServiceCode = &v
	return s
}

func (s *ExtractAddressRequest) SetText(v string) *ExtractAddressRequest {
	s.Text = &v
	return s
}

func (s *ExtractAddressRequest) SetDefaultProvince(v string) *ExtractAddressRequest {
	s.DefaultProvince = &v
	return s
}

func (s *ExtractAddressRequest) SetDefaultCity(v string) *ExtractAddressRequest {
	s.DefaultCity = &v
	return s
}

func (s *ExtractAddressRequest) SetDefaultDistrict(v string) *ExtractAddressRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *ExtractAddressRequest) SetAppKey(v string) *ExtractAddressRequest {
	s.AppKey = &v
	return s
}

type ExtractAddressResponse struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ExtractAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s ExtractAddressResponse) GoString() string {
	return s.String()
}

func (s *ExtractAddressResponse) SetData(v string) *ExtractAddressResponse {
	s.Data = &v
	return s
}

func (s *ExtractAddressResponse) SetRequestId(v string) *ExtractAddressResponse {
	s.RequestId = &v
	return s
}

type UpdateProjectRequest struct {
	ServiceCode *string                `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty" require:"true"`
	Parameters  map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty" require:"true"`
}

func (s UpdateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) SetServiceCode(v string) *UpdateProjectRequest {
	s.ServiceCode = &v
	return s
}

func (s *UpdateProjectRequest) SetParameters(v map[string]interface{}) *UpdateProjectRequest {
	s.Parameters = v
	return s
}

type UpdateProjectShrinkRequest struct {
	ServiceCode      *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty" require:"true"`
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty" require:"true"`
}

func (s UpdateProjectShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectShrinkRequest) SetServiceCode(v string) *UpdateProjectShrinkRequest {
	s.ServiceCode = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetParametersShrink(v string) *UpdateProjectShrinkRequest {
	s.ParametersShrink = &v
	return s
}

type UpdateProjectResponse struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponse) SetData(v string) *UpdateProjectResponse {
	s.Data = &v
	return s
}

func (s *UpdateProjectResponse) SetRequestId(v string) *UpdateProjectResponse {
	s.RequestId = &v
	return s
}

type CorrectAddressRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty" require:"true"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty" require:"true"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty" require:"true"`
}

func (s CorrectAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s CorrectAddressRequest) GoString() string {
	return s.String()
}

func (s *CorrectAddressRequest) SetServiceCode(v string) *CorrectAddressRequest {
	s.ServiceCode = &v
	return s
}

func (s *CorrectAddressRequest) SetText(v string) *CorrectAddressRequest {
	s.Text = &v
	return s
}

func (s *CorrectAddressRequest) SetDefaultProvince(v string) *CorrectAddressRequest {
	s.DefaultProvince = &v
	return s
}

func (s *CorrectAddressRequest) SetDefaultCity(v string) *CorrectAddressRequest {
	s.DefaultCity = &v
	return s
}

func (s *CorrectAddressRequest) SetDefaultDistrict(v string) *CorrectAddressRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *CorrectAddressRequest) SetAppKey(v string) *CorrectAddressRequest {
	s.AppKey = &v
	return s
}

type CorrectAddressResponse struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s CorrectAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s CorrectAddressResponse) GoString() string {
	return s.String()
}

func (s *CorrectAddressResponse) SetData(v string) *CorrectAddressResponse {
	s.Data = &v
	return s
}

func (s *CorrectAddressResponse) SetRequestId(v string) *CorrectAddressResponse {
	s.RequestId = &v
	return s
}

type GetAddressSimilarityRequest struct {
	ServiceCode     *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty" require:"true"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty" require:"true"`
	DefaultProvince *string `json:"DefaultProvince,omitempty" xml:"DefaultProvince,omitempty"`
	DefaultCity     *string `json:"DefaultCity,omitempty" xml:"DefaultCity,omitempty"`
	DefaultDistrict *string `json:"DefaultDistrict,omitempty" xml:"DefaultDistrict,omitempty"`
	AppKey          *string `json:"AppKey,omitempty" xml:"AppKey,omitempty" require:"true"`
}

func (s GetAddressSimilarityRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAddressSimilarityRequest) GoString() string {
	return s.String()
}

func (s *GetAddressSimilarityRequest) SetServiceCode(v string) *GetAddressSimilarityRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetAddressSimilarityRequest) SetText(v string) *GetAddressSimilarityRequest {
	s.Text = &v
	return s
}

func (s *GetAddressSimilarityRequest) SetDefaultProvince(v string) *GetAddressSimilarityRequest {
	s.DefaultProvince = &v
	return s
}

func (s *GetAddressSimilarityRequest) SetDefaultCity(v string) *GetAddressSimilarityRequest {
	s.DefaultCity = &v
	return s
}

func (s *GetAddressSimilarityRequest) SetDefaultDistrict(v string) *GetAddressSimilarityRequest {
	s.DefaultDistrict = &v
	return s
}

func (s *GetAddressSimilarityRequest) SetAppKey(v string) *GetAddressSimilarityRequest {
	s.AppKey = &v
	return s
}

type GetAddressSimilarityResponse struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s GetAddressSimilarityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAddressSimilarityResponse) GoString() string {
	return s.String()
}

func (s *GetAddressSimilarityResponse) SetData(v string) *GetAddressSimilarityResponse {
	s.Data = &v
	return s
}

func (s *GetAddressSimilarityResponse) SetRequestId(v string) *GetAddressSimilarityResponse {
	s.RequestId = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("address-purification"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetAddressGeocodeWithOptions(request *GetAddressGeocodeRequest, runtime *util.RuntimeOptions) (_result *GetAddressGeocodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetAddressGeocodeResponse{}
	_body, _err := client.DoRequest(tea.String("GetAddressGeocode"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-11-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAddressGeocode(request *GetAddressGeocodeRequest) (_result *GetAddressGeocodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAddressGeocodeResponse{}
	_body, _err := client.GetAddressGeocodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CompleteAddressWithOptions(request *CompleteAddressRequest, runtime *util.RuntimeOptions) (_result *CompleteAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CompleteAddressResponse{}
	_body, _err := client.DoRequest(tea.String("CompleteAddress"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-11-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CompleteAddress(request *CompleteAddressRequest) (_result *CompleteAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CompleteAddressResponse{}
	_body, _err := client.CompleteAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetZipcodeWithOptions(request *GetZipcodeRequest, runtime *util.RuntimeOptions) (_result *GetZipcodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetZipcodeResponse{}
	_body, _err := client.DoRequest(tea.String("GetZipcode"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-11-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetZipcode(request *GetZipcodeRequest) (_result *GetZipcodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetZipcodeResponse{}
	_body, _err := client.GetZipcodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExtractPhoneWithOptions(request *ExtractPhoneRequest, runtime *util.RuntimeOptions) (_result *ExtractPhoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ExtractPhoneResponse{}
	_body, _err := client.DoRequest(tea.String("ExtractPhone"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-11-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExtractPhone(request *ExtractPhoneRequest) (_result *ExtractPhoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExtractPhoneResponse{}
	_body, _err := client.ExtractPhoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExtractNameWithOptions(request *ExtractNameRequest, runtime *util.RuntimeOptions) (_result *ExtractNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ExtractNameResponse{}
	_body, _err := client.DoRequest(tea.String("ExtractName"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-11-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExtractName(request *ExtractNameRequest) (_result *ExtractNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExtractNameResponse{}
	_body, _err := client.ExtractNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAddressDivisionCodeWithOptions(request *GetAddressDivisionCodeRequest, runtime *util.RuntimeOptions) (_result *GetAddressDivisionCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetAddressDivisionCodeResponse{}
	_body, _err := client.DoRequest(tea.String("GetAddressDivisionCode"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-11-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAddressDivisionCode(request *GetAddressDivisionCodeRequest) (_result *GetAddressDivisionCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAddressDivisionCodeResponse{}
	_body, _err := client.GetAddressDivisionCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClassifyPOIWithOptions(request *ClassifyPOIRequest, runtime *util.RuntimeOptions) (_result *ClassifyPOIResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ClassifyPOIResponse{}
	_body, _err := client.DoRequest(tea.String("ClassifyPOI"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-11-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClassifyPOI(request *ClassifyPOIRequest) (_result *ClassifyPOIResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClassifyPOIResponse{}
	_body, _err := client.ClassifyPOIWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StructureAddressWithOptions(request *StructureAddressRequest, runtime *util.RuntimeOptions) (_result *StructureAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StructureAddressResponse{}
	_body, _err := client.DoRequest(tea.String("StructureAddress"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-11-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StructureAddress(request *StructureAddressRequest) (_result *StructureAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StructureAddressResponse{}
	_body, _err := client.StructureAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExtractAddressWithOptions(request *ExtractAddressRequest, runtime *util.RuntimeOptions) (_result *ExtractAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ExtractAddressResponse{}
	_body, _err := client.DoRequest(tea.String("ExtractAddress"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-11-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExtractAddress(request *ExtractAddressRequest) (_result *ExtractAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExtractAddressResponse{}
	_body, _err := client.ExtractAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProjectWithOptions(tmp *UpdateProjectRequest, runtime *util.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
	_err = util.ValidateModel(tmp)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateProjectShrinkRequest{}
	rpcutil.Convert(tmp, request)
	if !tea.BoolValue(util.IsUnset(tmp.Parameters)) {
		request.ParametersShrink = util.ToJSONString(tmp.Parameters)
	}

	_result = &UpdateProjectResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateProject"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-11-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProject(request *UpdateProjectRequest) (_result *UpdateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateProjectResponse{}
	_body, _err := client.UpdateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CorrectAddressWithOptions(request *CorrectAddressRequest, runtime *util.RuntimeOptions) (_result *CorrectAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CorrectAddressResponse{}
	_body, _err := client.DoRequest(tea.String("CorrectAddress"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-11-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CorrectAddress(request *CorrectAddressRequest) (_result *CorrectAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CorrectAddressResponse{}
	_body, _err := client.CorrectAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAddressSimilarityWithOptions(request *GetAddressSimilarityRequest, runtime *util.RuntimeOptions) (_result *GetAddressSimilarityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetAddressSimilarityResponse{}
	_body, _err := client.DoRequest(tea.String("GetAddressSimilarity"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-11-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAddressSimilarity(request *GetAddressSimilarityRequest) (_result *GetAddressSimilarityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAddressSimilarityResponse{}
	_body, _err := client.GetAddressSimilarityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
